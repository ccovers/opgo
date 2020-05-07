import MySQLdb
import pandas.io.sql as psql
import pandas as pd
import numpy as np
import google.auth
from google.cloud import bigquery

from google.oauth2 import service_account
import difflib
import os
import json
import time
start_time = time.time()

project_id = 'axinan-data'

# service account to read and write to axinan-data: Bigquery
# for large datasets and complicated queries, performance of MySQL is a bottleneck
# thus read data to Bigquery first
# also accessing data from Bigquery is much easier, without proxy requirements
credentials = service_account.Credentials.from_service_account_file(
    './credentials/axinan-data-edba77dccd9a.json')

client = bigquery.Client(credentials=credentials, project=project_id)

# function to get data from newtond_shopee_xb
# user and passwd as input not recommended, need a recommendation how to improve
def get_data_soure(query,host='127.0.0.1', user='jack.xia', passwd='fei5xie8ShoNgi1Roh0g', db='newtond_shopee_xb',port=3308):
	db=MySQLdb.connect(host=host, user=user, passwd=passwd, db=db,port=port)
	table = psql.read_sql(query, con=db)
	db.close()
	return table

# expand all json nested fields to fields in table then save to bigquery
# note here there can be multiple items in the same claims
# here just take the first item, multi-items to process separately
# dont expand multi items directly here, will result problems in other non-json fields
def update_claim():
    ##Step 1 update claim table
    sql = """
        SELECT
        id,
        status,
        claimType,
        refundStatus,
        JSON_LENGTH(details, '$.Raw.items') as total_items,
        json_unquote(json_extract(details, '$.Raw.items[0].reason')) as claim_reason,
        json_unquote(json_extract(details, '$.Raw.items[0]."3plProof".type')) as Proof_type,
        json_unquote(json_extract(details, '$.Raw.items[0]."3plProof".reason')) as Proof_reason,
        json_unquote(json_extract(details, '$.Raw.items[0]."3plProof".refundStatus')) as Proof_refundStatus,
        json_unquote(json_extract(details, '$.Raw.items[0]."3plProof".shippingStatus')) as Proof_shippingStatus,
        json_unquote(json_extract(details, '$.Raw.items[0]."3plProof".statusUpdatedAt')) as Proof_statusUpdatedAt,
        json_unquote(json_extract(details, '$.Raw.items[0]."3plProof".returnRequestedAt')) as Proof_returnRequestedAt,
        json_unquote(json_extract(details, '$.Raw.items[0].lineitem.id')) as item_id,
        json_unquote(json_extract(details, '$.Raw.items[0].lineitem.qty')) as item_qty,
        json_unquote(json_extract(details, '$.Raw.items[0].lineitem.sku')) as item_sku,
        json_unquote(json_extract(details, '$.Raw.items[0].lineitem.name')) as item_name,
        json_unquote(json_extract(details, '$.Raw.items[0].lineitem.price')) as item_price,
        json_unquote(json_extract(details, '$.Raw.items[0].lineitem.productLink')) as item_link,
        json_unquote(json_extract(details, '$.Raw.items[0].lineitem.rackingNumber')) as item_trackingNumber,
        json_unquote(json_extract(details, '$.Raw.items[0].lineitem.packageTotalAmount')) as item_packageTotalAmount,
        json_unquote(json_extract(details, '$.Raw.items[0].chronology')) as chronology,
        json_unquote(json_extract(details, '$.Raw.items[0].reasonCode')) as reasonCode,
        json_unquote(json_extract(details, '$.Raw.items[0].claimAmount')) as item_claimAmount,
        json_unquote(json_extract(details, '$.Raw.items[0].shippingFee')) as shippingFee,
        json_unquote(json_extract(details, '$.Raw.items[0].refundAmount')) as item_refundAmount,
        json_unquote(json_extract(details, '$.Raw.items[0].shippingNumber')) as shippingNumber,
        json_unquote(json_extract(details, '$.Raw.cancelReason')) as cancelReason,
        json_unquote(json_extract(details, '$.Raw.rejectReason')) as rejectReason,
        json_unquote(json_extract(details, '$.Raw.claimMaxAmount')) as claimMaxAmount,
        json_unquote(json_extract(details, '$.Raw.refundAmountCNY')) as refundAmountCNY,
        json_unquote(json_extract(details, '$.Raw.refundAmountUSD')) as refundAmountUSD,
        json_unquote(json_extract(details, '$.Raw.claimHistoryList')) as claimHistoryList,
        json_unquote(json_extract(details, '$.Raw.refundAmountHistory.reason')) as refundAmountHistory_reason,
        json_unquote(json_extract(details, '$.Raw.refundAmountHistory.refundAmount')) as refundAmountHistory_refundAmount,
        json_unquote(json_extract(details, '$.Raw.needManualReimbursement')) as needManualReimbursement,
        claimAmount,
        refundAmount,
        currency,
        axinanApprovedAt,
        createdAt,
        updatedAt,
        processingAt,
        filedAt,
        extOrderId,
        uploadedAt,
        refundPendingAt,
        lastAxinanOperator,
        policyId,
        shopId,
        platform,
        orderCreatedAt

        FROM newtond_shopee_xb.Claim
    """

    # load all data into pandas dataframe
    # because lazada and shopee xb datasets are small so can be done this way
    # larger datasets need a pipeline to transfer from MySQL to Bigquery directly
    claim = get_data_soure(query=sql, db='newtond_shopee_xb')

    print("Uploading claim data to bigquery")

    # upload dataframe as new table into Bigquery
    claim.to_gbq("shopee_XB_Claim.claim",
                      project_id=project_id,
                      if_exists="replace",
                      credentials=credentials)
    print("Successfully updated claim")

    # for some claims, there are multiple items in the same claim. Iterate over the items to record all item details
    sql = """
        SELECT
            max(JSON_LENGTH(details, '$.Raw.items')) as max_items
        FROM newtond_shopee_xb.Claim
    """

    # get the maximum possible items and iterate
    max = get_data_soure(query=sql, db='newtond_shopee_xb')
    max_iter = max['max_items'][0]

    # expand data in SQL instead of expanding in Python. May not be the most efficient but works
    # ignore fields with claim level numbers, only show item level info
    claim_iter = pd.DataFrame()
    for i in range(max_iter):
        print("Merging items in claim :", i)
        sql = """
                SELECT
                id,
                status,
                claimType,
                refundStatus,
                JSON_LENGTH(details, '$.Raw.items') as total_items,
                json_unquote(json_extract(details, '$.Raw.items[%i].lineitem.id')) as item_id,
                json_unquote(json_extract(details, '$.Raw.items[%i].lineitem.qty')) as item_qty,
                json_unquote(json_extract(details, '$.Raw.items[%i].lineitem.sku')) as item_sku,
                json_unquote(json_extract(details, '$.Raw.items[%i].lineitem.name')) as item_name,
                json_unquote(json_extract(details, '$.Raw.items[%i].lineitem.price')) as item_price,
                json_unquote(json_extract(details, '$.Raw.items[%i].lineitem.productLink')) as item_link,
                json_unquote(json_extract(details, '$.Raw.items[%i].chronology')) as chronology,
                json_unquote(json_extract(details, '$.Raw.items[%i].reasonCode')) as reasonCode,
                json_unquote(json_extract(details, '$.Raw.items[%i].claimAmount')) as item_claimAmount,
                json_unquote(json_extract(details, '$.Raw.items[%i].shippingFee')) as shippingFee,
                json_unquote(json_extract(details, '$.Raw.items[%i].refundAmount')) as item_refundAmount,
                json_unquote(json_extract(details, '$.Raw.items[%i].shippingNumber')) as shippingNumber,
                extOrderId,
                policyId,
                shopId

                FROM newtond_shopee_xb.Claim
                where JSON_LENGTH(details, '$.Raw.items') >= %i
            """

        # the last >= means if there are 10 items in a claim, this item is iterated 10 times
        sql = sql % (i, i, i, i, i, i, i, i, i, i, i, i, i + 1)
        claim_temp = get_data_soure(query=sql, db='newtond_shopee_xb')

        # append to dataframe, first time append to empty will initialize schema
        claim_iter = pd.concat([claim_iter, claim_temp], axis=0, ignore_index=True)

    claim_iter.to_gbq("shopee_XB_Claim.claim_expanditem",
                      project_id=project_id,
                      if_exists="replace",
                      credentials=credentials)
    print("Successfully updated claim with expanded items")



# we need item, policy, order level information to verify the claim data
# load all of them to Bigquery
# this step may take some time
def update_item_policy_order():
    sql = """
        SELECT
        *
        FROM newtond_shopee_xb.Item
    """
    claim = get_data_soure(query=sql, db='newtond_shopee_xb')

    print("Uploading item data to bigquery")
    claim.to_gbq("shopee_XB_Claim.item",
                 project_id=project_id,
                 if_exists="replace",
                 credentials=credentials)
    print("Successfully updated item")

    sql = """
        SELECT
        *
        FROM newtond_shopee_xb.Policy
    """
    claim = get_data_soure(query=sql, db='newtond_shopee_xb')

    print("Uploading Policy data to bigquery")
    claim.to_gbq("shopee_XB_Claim.Policy",
                 project_id=project_id,
                 if_exists="replace",
                 credentials=credentials)
    print("Successfully updated Policy")

    sql = """
        SELECT
        *
        FROM newtond_shopee_xb.Order
    """
    claim = get_data_soure(query=sql, db='newtond_shopee_xb')

    print("Uploading item data to bigquery")
    claim.to_gbq("shopee_XB_Claim.Order",
                 project_id=project_id,
                 if_exists="replace",
                 credentials=credentials)
    print("Successfully updated Order")

# similarly, update shop
# this step is faster
def update_shop():
    sql = """
        SELECT
        *
        FROM newtond_shopee_xb.Shop
    """
    claim = get_data_soure(query=sql, db='newtond_shopee_xb')

    print("Uploading shop data to bigquery")
    claim.to_gbq("shopee_XB_Claim.shop",
                 project_id=project_id,
                 if_exists="replace",
                 credentials=credentials)
    print("Successfully updated shop")

#similarly, update invoice
#this is as slow as updating policy, there is an Invoice for each Policy
def update_invoice():
    sql = """
        SELECT
        *
        FROM newtond_shopee_xb.InvoiceV2
    """
    claim = get_data_soure(query=sql, db='newtond_shopee_xb')

    print("Uploading InvoiceV2 data to bigquery")
    claim.to_gbq("shopee_XB_Claim.InvoiceV2",
                 project_id=project_id,
                 if_exists="replace",
                 credentials=credentials)
    print("Successfully updated InvoiceV2")

    sql = """
            SELECT
            *
            FROM newtond_shopee_xb.InvoiceBatchV2
        """
    claim = get_data_soure(query=sql, db='newtond_shopee_xb')

    print("Uploading InvoiceBatchV2 data to bigquery")
    claim.to_gbq("shopee_XB_Claim.InvoiceBatchV2",
                 project_id=project_id,
                 if_exists="replace",
                 credentials=credentials)
    print("Successfully updated InvoiceBatchV2")

# check all potential problems and inconsistencies in the claim table
# to add if some other problems are found
# to edit depends on feedback
def check_claim():
    #for high loss ratio merchants, to check if they have been freezed/unregistered/freezed
    # or premium adjustment has been implemented
    df_freezeClaim = client.query('''
        SELECT
          c.ownerMail,
          c.claimUSD_approved,
          p.premiumAmount_uploaded,
          s.deletedAt,
          s.unRegisteredAt,
          s.freezedAt
        FROM
        (
          SELECT
            s.ownerMail,
            count(claimAmount) as claimCount,
            sum(cast(claimAmount as int64)/100 * e.rate * cast(s.plan as int64)/100) as claimUSD_uploaded,
            sum(cast(claimAmount as int64)/100 * e.rate * cast(s.plan as int64)/100 *
            CASE WHEN c.status = "approved" then 1 else 0 END) as claimUSD_approved
          FROM
          `axinan-data.shopee_XB_Claim.claim` as c
          left join
          `axinan-data.shopee_XB_Claim.exchange_rate` as e
          on c.currency = e.currency
          left join
          `axinan-data.shopee_XB_Claim.shop` as s
          on c.shopId = s.id
          where s.deletedAt is null
          and TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), cast(c.uploadedAt as timestamp), DAY) <= 14
          group by s.ownerMail
        ) as c
        left join
        (
          SELECT
            s.ownerMail,
            count(premiumAmount) as PolicyCount,
            sum(cast(premiumAmount as int64)/100 * e.rate * cast(s.plan as int64)/100) as premiumAmount_uploaded
          FROM
          `axinan-data.shopee_XB_Claim.Policy` as p
          left join
          `axinan-data.shopee_XB_Claim.exchange_rate` as e
          on p.currency = e.currency
          left join
          `axinan-data.shopee_XB_Claim.shop` as s
          on p.shopId = s.id
          where s.deletedAt is null and p.deletedAt is null
          and TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), cast(p.createdAt as timestamp), DAY) <= 14
          group by s.ownerMail
        ) as p
        on c.ownerMail = p.ownerMail
        left join
        (
          SELECT
            ownerMail,
            deletedAt,
            max(unRegisteredAt) as unRegisteredAt,
            max(freezedAt) as freezedAt
          FROM `axinan-data.shopee_XB_Claim.shop`
          group by ownerMail, deletedAt
        ) as s
        on c.ownerMail = s.ownerMail

        where claimUSD_approved - premiumAmount_uploaded > 100 and claimUSD_approved/premiumAmount_uploaded > 1.5
        and   s.deletedAt is null and s.unRegisteredAt is null and s.freezedAt is null
        order by claimUSD_uploaded desc
                                    ''').to_dataframe()

    # if query does not produce anything then test PASS
    if df_freezeClaim.empty:
        print("Pass: no high claim sellers, OK to generate claim invoice")
    # otherwise, send the csv file to person-in-charge to verify
    # csv starts with person-in-charge
    else:
        df_freezeClaim.to_csv("Commercial_freezeClaim.csv")
        print("Following sellers with high claims in last 14 days, consider freeze claim invoice : Commercial_freezeClaim.csv")
        print(df_freezeClaim)

    # check there is no new reason-code combination
    # there was a big table from shopee to describe all possible combinations
    # to avoid mis-understanding, we just highlight all new unverified combinations
    # reasoncode table is generated separately
    df_reasonCode = client.query('''
            SELECT
            c.reasonCode,
            c.claim_reason,
            r.claim_reason as expected_reasoncode
            FROM
            `axinan-data.shopee_XB_Claim.claim` as c
            left join
            `axinan-data.shopee_XB_Claim.reasoncode` as r
            on c.reasonCode = r.reasonCode
            where r.claim_reason is NULL or c.claim_reason != r.claim_reason
            ''').to_dataframe()

    if df_reasonCode.empty:
        print("Pass: reasoncode check")
    else:
        df_reasonCode.to_csv("reasonCode.csv")
        print("To check unexpected reasonCode: reasonCode.csv")

    # same for claim proof type
    df_claim_proof_Type = client.query('''
                select
                c.*
                FROM
                `axinan-data.shopee_XB_Claim.claim` as c
                left join
                `axinan-data.shopee_XB_Claim.claim_proof_Type` as t
                on c.claimType = t.claimType and c.Proof_type = t.Proof_type
                where t.Proof_type is NULL

    ''').to_dataframe()

    if df_claim_proof_Type.empty:
        print("Pass: unexpected claim or proof type check")
    else:
        df_claim_proof_Type.to_csv("claim_proof_Type.csv")
        print("Check unexpected claim or proof type: claim_proof_Type.csv")

    # to check that 7 days return is really submitted within 7 days
    # and we assume cannot be submitted on the first day due to time taken for delivery
    df_claimType7daysFree = client.query('''
            SELECT
            id as claimid,
            policyid,
            claimType,
            orderCreatedAt,
            Proof_statusUpdatedAt,
            TIMESTAMP_DIFF(cast(Proof_statusUpdatedAt as timestamp), orderCreatedAt, DAY) as submit_claim_days,
            uploadedAt,
            TIMESTAMP_DIFF(uploadedAt, cast(Proof_statusUpdatedAt as timestamp), DAY) as upload_claim_days,

            FROM `axinan-data.shopee_XB_Claim.claim`
            where claimType = '7dayFreeReturn'
            and ( TIMESTAMP_DIFF(cast(Proof_statusUpdatedAt as timestamp), orderCreatedAt, DAY) < 2 or TIMESTAMP_DIFF(cast(Proof_statusUpdatedAt as timestamp), orderCreatedAt, DAY) > 7)
            ''').to_dataframe()

    if df_claimType7daysFree.empty:
        print("Pass: claimType 7daysFree check")
    else:
        df_claimType7daysFree.to_csv("df_claimType7daysFree.csv")
        print("To check within 1 day or more than 7 days 7dayFree return: df_claimType7daysFree.csv")

    # for damage, there is a reason usd 50+ and usd 50-
    # here we check if the item is really above or below 50
    # use a fixed exchange rate and allow $5 difference due to exchange rate
    df_claimTypedamage = client.query('''
        SELECT
            i.sku,
            i.priceUSD,
            c.*

        FROM
        `axinan-data.shopee_XB_Claim.claim` as c
        left join
        (select
            policyId,
            sku,
            cast(json_extract(details,'$.Raw.priceUSD') as int64)/100 as priceUSD
            from  `axinan-data.shopee_XB_Claim.item` ) as i
        on c.policyId = i.policyId

        where claimType = 'damage'
        and not (Proof_type = 'damaged (usd 50+)' and priceUSD > 45)
        and not (Proof_type = 'damaged (usd 50-)' and priceUSD < 55)
        ''').to_dataframe()

    if df_claimTypedamage.empty:
        print("Pass: claimType damage price consistency check")
    else:
        df_claimTypedamage.to_csv("Tech_claimTypedamage_price.csv")
        print("To check damage with invalid proof or price inconsistent with +- 50 USD: Tech_claimTypedamage_price.csv")

    # for damage claims there is a $50 USD excess
    # here we check excess is deducted
    # use fixed exchange rate and allow $5 difference
    df_claimTypedamage_refund = client.query('''
            SELECT
                i.sku,
                i.priceUSD,
                c.refundAmount * e.rate/100 as refund_amount,
                c.*

            FROM
            `axinan-data.shopee_XB_Claim.claim` as c
            left join
            (select policyId,sku, cast(json_extract(details,'$.Raw.priceUSD') as int64)/100 as priceUSD from  `axinan-data.shopee_XB_Claim.item` )as i
            on c.policyId = i.policyId
            left join
            `axinan-data.shopee_XB_Claim.exchange_rate` as e
            on e.currency = c.currency
            where claimType = 'damage'
              and ((Proof_type = 'damaged (usd 50-)' and c.refundAmount >0) or (Proof_type = 'damaged (usd 50+)' and c.refundAmount * e.rate/100 >i.priceUSD - 40))
            #note that exchange above is not taken from the transaction date, so we allow more tolerance
            ''').to_dataframe()

    if df_claimTypedamage_refund.empty:
        print("Pass: claimType damage price consistency check")
    else:
        df_claimTypedamage_refund.to_csv("Tech_claimTypedamage_refund.csv")
        print("To check damage refund amount did not minus the excess of USD 50: Tech_claimTypedamage_refund.csv")

    # To check delay claim submission is X days after create date,
    # 7 days for SG, TW; 12 days for VN and 15 days for TH, PH and ID
    # should reject the claims within these days automatically
    df_claimTypedelay= client.query('''
        SELECT
            id as check_claimid,
            policyid as check_policyid,
            claimType as check_claimType,
            orderCreatedAt as check_orderCreatedAt,
            Proof_statusUpdatedAt as check_Proof_statusUpdatedAt,
            submit_claim_days as delay_submitted_days,
            uploadedAt as check_uploadedAt,
            TIMESTAMP_DIFF(uploadedAt, cast(Proof_statusUpdatedAt as timestamp), DAY) as upload_claim_days,
            *

        FROM
        (select
          *,
          TIMESTAMP_DIFF(cast(Proof_statusUpdatedAt as timestamp), orderCreatedAt, DAY) as submit_claim_days
         from
          `axinan-data.shopee_XB_Claim.claim`) as c
        where
          claimType = 'delay' and
          status != 'rejected' and # should reject these claims
          (
            (submit_claim_days < 8 and currency in ('SGD', 'TWD')) or # 7 days for SG and TW
            (submit_claim_days < 12 and currency in ('VND')) or #11 days for VN
            (submit_claim_days < 15 and currency in ('THB', 'PHP','IDR')) #15 days for TH, PH, ID
          )
        order by c.submit_claim_days asc
        ''').to_dataframe()

    if df_claimTypedelay.empty:
        print("Pass: claimType delay check")
    else:
        df_claimTypedelay.to_csv("Data_claimTypeddelay.csv")
        print("To check claimed delay before X days: Data_claimTypeddelay.csv")

    # To check that no merchant is mis-selling products with descriptions difference
    # Accept that some customers will make such claims naturally
    # The threshold now is $50 USD or 0.5% of items
    df_claimTypedescription_shops = client.query('''
            SELECT
                c.shopid,
                total_USD,
                description_difference_counts,
                item_count
                FROM
                (select
                    shopid,
                    sum(c.claimAmount * e.rate /100) as total_USD,
                    count(shopid) as description_difference_counts
                  from `axinan-data.shopee_XB_Claim.claim` as c
                  left join `axinan-data.shopee_XB_Claim.exchange_rate` as e
                  on c.currency = e.currency
                  where claimType = 'descriptionDifference'
                  group by shopid) as c
                left join
                (select shopid, count(shopid) as item_count from `axinan-data.shopee_XB_Claim.item` group by shopid) as i
                on c.shopid = i.shopid
                where total_USD > 50 #if there a small items with description difference, it is accepted
                  or description_difference_counts/item_count > 0.005 # if there is a small amount of description difference claims, it is accepted
            order by description_difference_counts desc
        ''').to_dataframe()

    if df_claimTypedescription_shops.empty:
        print("Pass: claimType description difference check")
    else:
        print("Following stores have description difference claim. To manual review if number is high.")
        print(df_claimTypedescription.head(100))

        df_claimTypedescription = client.query('''
                    SELECT
                    i.sku,
                    cast(json_extract(details,'$.Raw.priceUSD') as int64)/100 as priceUSD,
                    c.*

                    FROM
                    `axinan-data.shopee_XB_Claim.claim` as c
                    left join
                    `axinan-data.shopee_XB_Claim.item` as i
                    on c.policyId = i.policyId

                    where claimType = 'descriptionDifference'
                ''').to_dataframe()
        df_claimTypedescription_shops.to_csv("Ops_claimTypedescription_shops.csv")
        print("To check following shops with significant description difference claims: Ops_claimTypedescription_shops.csv")
        df_claimTypedescription.to_csv("Ops_claimTypedescription.csv")
        print("To check description details of the items: Ops_claimTypedescription.csv")

    # same as damage, check lost USD 50+ is above 50 usd
    df_lost50_price= client.query('''
        SELECT
            i.sku,
            i.priceUSD,
            c.*
        FROM
        `axinan-data.shopee_XB_Claim.claim` as c
        left join
        (select policyId,sku, cast(json_extract(details,'$.Raw.priceUSD') as int64)/100 as priceUSD from  `axinan-data.shopee_XB_Claim.item` )as i
        on c.policyId = i.policyId

        where claimType = 'lost' and Proof_type = 'lost (usd 50+)' and priceUSD <55
        ''').to_dataframe()

    if df_lost50_price.empty:
        print("Pass: lost above 50 USD product price check")
    else:
        df_lost50_price.to_csv("Tech_lost50_price.csv")
        print("To check lost with usd 50+ but price less than 50: Tech_lost50_price.csv")

    # check refund amount is after excess
    df_lost50_refund = client.query('''
        SELECT
            i.sku,
            i.priceUSD,
            c.refundAmount * e.rate/100 as refund_amount,
            c.*
        FROM
        `axinan-data.shopee_XB_Claim.claim` as c
        left join
        (select policyId,sku, cast(json_extract(details,'$.Raw.priceUSD') as int64)/100 as priceUSD from  `axinan-data.shopee_XB_Claim.item` )as i
        on c.policyId = i.policyId
        left join
        `axinan-data.shopee_XB_Claim.exchange_rate` as e
        on e.currency = c.currency
        where claimType = 'lost'
          and Proof_type = 'lost (usd 50+)' and c.refundAmount * e.rate/100 >i.priceUSD - 40
        #note that exchange above is not taken from the transaction date, so we allow more tolerance
        ''').to_dataframe()

    if df_lost50_refund.empty:
        print("Pass: lost above 50 USD refund amount check")
    else:
        df_lost50_refund.to_csv("Tech_lost50_refund.csv")
        print("To check lost with usd 50+ but price less than 50: Tech_lost50_refund.csv")

    # for lost and exclude insurance, the claims should not be approved
    df_lost_exclude = client.query('''
        SELECT
            i.sku,
            i.priceUSD,
            c.*
        FROM
        `axinan-data.shopee_XB_Claim.claim` as c
        left join
        (select policyId,sku, cast(json_extract(details,'$.Raw.priceUSD') as int64)/100 as priceUSD from  `axinan-data.shopee_XB_Claim.item` )as i
        on c.policyId = i.policyId

        where claimType = 'lost' and Proof_type = 'exclude_insurance' and status = 'approved'
            ''').to_dataframe()

    if df_lost_exclude.empty:
        print("Pass: lost exclude_insurance_check check")
    else:
        df_lost_exclude.to_csv("Tech_lost_exclude.csv")
        print("To check lost exclude_insurance but approved: Tech_lost_exclude.csv")

    # for lost, the lost status should only be updated after 40 days
    df_lost_40days = client.query('''
                SELECT
                id as claimid,
                policyid,
                claimType,
                orderCreatedAt,
                Proof_statusUpdatedAt,
                TIMESTAMP_DIFF(cast(Proof_statusUpdatedAt as timestamp), orderCreatedAt, DAY) as submit_claim_days,
                uploadedAt,
                TIMESTAMP_DIFF(uploadedAt, cast(Proof_statusUpdatedAt as timestamp), DAY) as upload_claim_days,

                FROM `axinan-data.shopee_XB_Claim.claim`
                where claimType = 'lost' and TIMESTAMP_DIFF(uploadedAt,cast(Proof_statusUpdatedAt as timestamp), DAY) < 40
                order by submit_claim_days asc
            ''').to_dataframe()

    if df_lost_40days.empty:
        print("Pass: lost only report after 40 days check")
    else:
        df_lost_40days.to_csv("df_lost_40days.csv")
        print("To check lost claim that are submitted within 40 days of order: df_lost_40days.csv")

    # check no unexpected lost status
    # may be duplicated from the all reason check above
    df_lost_unexpected = client.query('''
                SELECT
                i.sku,
                i.priceUSD,
                c.*
                FROM
                `axinan-data.shopee_XB_Claim.claim` as c
                left join
                (select policyId,sku, cast(json_extract(details,'$.Raw.priceUSD') as int64)/100 as priceUSD from  `axinan-data.shopee_XB_Claim.item` )as i
                on c.policyId = i.policyId

                where claimType = 'lost' and not (Proof_type = 'exclude_insurance' or Proof_type = 'lost (usd 50+)')
            ''').to_dataframe()

    if df_lost_unexpected.empty:
        print("Pass: lost proof type check")
    else:
        df_lost_unexpected.to_csv("lost_unexpected.csv")
        print("To check unexpected lost proof: lost_unexpected.csv")

    # check missing or wrong item after 7 days is really after 7 days
    # allow a few more days for delivery time
    df_morr7days = client.query('''
                SELECT
                i.sku,
                i.priceUSD,
                TIMESTAMP_DIFF(cast(Proof_statusUpdatedAt as timestamp), orderCreatedAt, DAY) as submit_claim_days,
                c.*


                FROM
                `axinan-data.shopee_XB_Claim.claim` as c
                left join
                (select policyId,sku, cast(json_extract(details,'$.Raw.priceUSD') as int64)/100 as priceUSD from  `axinan-data.shopee_XB_Claim.item` )as i
                on c.policyId = i.policyId

                where claimType = 'missingOrWrongItem' and Proof_type = '7 days no reason to return'
                # allow some days for delivery so > 10
                and TIMESTAMP_DIFF(cast(Proof_statusUpdatedAt as timestamp), orderCreatedAt, DAY) > 10
            ''').to_dataframe()

    if df_morr7days.empty:
        print("Pass: missing or wrong 7 days return check")
    else:
        df_morr7days.to_csv("lost_morr7days.csv")
        print("To check returned after 10 days of order for 7 days no reason to return: df_morr7days.csv")

    # check claim refund status combination
    df_claim_refund_status = client.query('''
                SELECT
                c.*

                FROM
                `axinan-data.shopee_XB_Claim.claim` as c
                left join
                `axinan-data.shopee_XB_Claim.item` as i
                on c.policyId = i.policyId
                left join
                `axinan-data.shopee_XB_Claim.claim_refund_status` as s
                on c.status= s.status and c.refundStatus = s.refundStatus

                where s.refundStatus is NULL
            ''').to_dataframe()

    if df_claim_refund_status.empty:
        print("Pass: no unexpected claim - refund status")
    else:
        df_claim_refund_status.to_csv("claim_refund_status.csv")
        print("To check unexpected claim status and refund status combinations: df_claim_refund_status.csv")

    # To check refund is paid within 7 days
    # Due to transaction cost, we do not pay if the claim invoice amount is small, here limited at $20 USD
    # If the shop level table shows a significant amount, go to check the details table
    df_refund_pending_days_shop = client.query('''
            SELECT
            *
            FROM
            (
              SELECT
                s.ownerMail,
                TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), cast(c.refundPendingAt as timestamp), DAY) as refund_pending_days,
                count(*) as claims,
                sum(cast(json_query(i.details,'$.Raw.priceUSD') as int64)/100) as refund_amount
              FROM
              `axinan-data.shopee_XB_Claim.claim` as c
              left join
              `axinan-data.shopee_XB_Claim.item` as i
              on c.policyId = i.policyId
              left join
              `axinan-data.shopee_XB_Claim.shop` as s
              on i.shopId = s.Id
              left join
              (select  i.policyId, b.status as premium_status
                from `axinan-data.shopee_XB_Claim.InvoiceV2` as i
                left join
                `axinan-data.shopee_XB_Claim.InvoiceBatchV2` as b
                on i.batchId = b.id
                where i.type = 'premium' and i.deletedAt is null and b.deletedAt is null) as pi
              on i.policyId = pi.policyId

              where c.status = "approved" and c.refundStatus = "pending"
              and s.deletedAt is null and s.unRegisteredAt is null and s.freezedAt is null
              and TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), cast(c.refundPendingAt as timestamp), DAY) > 7

              group by ownerMail, refund_pending_days
            )
            where refund_amount > 20 # due to transaction cost, we do not pay small claims, limit by 20
            order by refund_pending_days desc
        ''').to_dataframe()

    if df_refund_pending_days_shop.empty:
        print("Pass: refund pending days difference check")
    else:
        print("Following stores have refund pending more than 7 days. Shop is not freezed, unregistered, deleted. Premium is paid.")
        print(df_refund_pending_days_shop.head(100))
        df_refund_pending_days_shop.to_csv("Ops_refund_pending_days_shop.csv")
        print("To check these shops with refund pending: Ops_refund_pending_days_shop.csv")

        df_refund_pending_days = client.query('''
                    SELECT
                        s.ownerMail,
                        s.deletedAt as shop_deletedAt,
                        s.unRegisteredAt as shop_unregisteredAt,
                        s.freezedAt as shop_freezdAt,
                        pi.premium_status,
                        TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), cast(c.refundPendingAt as timestamp), DAY) as refund_pending_days,
                        c.*

                    FROM
                    `axinan-data.shopee_XB_Claim.claim` as c
                    left join
                    `axinan-data.shopee_XB_Claim.item` as i
                    on c.policyId = i.policyId
                    left join
                    `axinan-data.shopee_XB_Claim.shop` as s
                    on i.shopId = s.Id
                    left join
                    (select  i.policyId, b.status as premium_status
                      from `axinan-data.shopee_XB_Claim.InvoiceV2` as i
                      left join
                      `axinan-data.shopee_XB_Claim.InvoiceBatchV2` as b
                      on i.batchId = b.id
                      where i.type = 'premium' and i.deletedAt is null and b.deletedAt is null) as pi
                    on i.policyId = pi.policyId

                    where c.status = "approved" and c.refundStatus = "pending"
                    and s.deletedAt is null and s.unRegisteredAt is null and s.freezedAt is null
                    and TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), cast(c.refundPendingAt as timestamp), DAY) > 7
                ''').to_dataframe()
        df_refund_pending_days.to_csv("Ops_refund_pending_days.csv")
        print("To check refund pending details, only focus on the shops with high amount: Ops_refund_pending_days.csv")

    # all unknown status should be processed within 7 days
    # so we can generate invoice every week
    # or use some other indicator to flag
    df_unknown_refund_status = client.query('''
            SELECT
                s.ownerMail,
                TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), cast(c.uploadedAt as timestamp), DAY) as uploaded_days,
                count(*) as claims
            FROM
            `axinan-data.shopee_XB_Claim.claim` as c
            left join
            `axinan-data.shopee_XB_Claim.item` as i
            on c.policyId = i.policyId
            left join
            `axinan-data.shopee_XB_Claim.shop` as s
            on i.shopId = s.Id
            left join
            (select  i.policyId, b.status as premium_status
              from `axinan-data.shopee_XB_Claim.InvoiceV2` as i
              left join
              `axinan-data.shopee_XB_Claim.InvoiceBatchV2` as b
              on i.batchId = b.id
              where i.type = 'premium' and i.deletedAt is null and b.deletedAt is null) as pi
            on i.policyId = pi.policyId

            where c.refundStatus = "unknown" and c.status != "rejected"
            and s.deletedAt is null and s.unRegisteredAt is null and s.freezedAt is null
            and TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), cast(c.uploadedAt as timestamp), DAY) > 7
            group by ownerMail, uploaded_days
            order by uploaded_days desc
        ''').to_dataframe()

    if df_unknown_refund_status.empty:
        print("Pass: refund unknown status check")
    else:
        print(
            "Following stores have unnknown refund more than 7 days. Shop is not freezed, unregistered, deleted. Premium is paid.")
        print(df_unknown_refund_status.head(100))

        df_unknown_refund_status= client.query('''
                    SELECT
                        s.ownerMail,
                        s.deletedAt as shop_deletedAt,
                        s.unRegisteredAt as shop_unregisteredAt,
                        s.freezedAt as shop_freezdAt,
                        pi.premium_status,
                        TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), cast(c.uploadedAt as timestamp), DAY) as uploaded_days,
                        c.*
                    FROM
                    `axinan-data.shopee_XB_Claim.claim` as c
                    left join
                    `axinan-data.shopee_XB_Claim.item` as i
                    on c.policyId = i.policyId
                    left join
                    `axinan-data.shopee_XB_Claim.shop` as s
                    on i.shopId = s.Id
                    left join
                    (select  i.policyId, b.status as premium_status
                      from `axinan-data.shopee_XB_Claim.InvoiceV2` as i
                      left join
                      `axinan-data.shopee_XB_Claim.InvoiceBatchV2` as b
                      on i.batchId = b.id
                      where i.type = 'premium' and i.deletedAt is null and b.deletedAt is null) as pi
                    on i.policyId = pi.policyId

                    where c.refundStatus = "unknown" and c.status != "rejected"
                    and s.deletedAt is null and s.unRegisteredAt is null and s.freezedAt is null
                    and TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), cast(c.uploadedAt as timestamp), DAY) > 7
                ''').to_dataframe()
        df_unknown_refund_status.to_csv("Tech_unknown_refund_status.csv")
        print("To check unknown refund details not updated with 7 days: Tech_unknown_refund_status.csv")

    # check that for single item claim, the claim amount is no more than the item amount
    # if shipping fee is included, need to get it from somewhere
    df_singleitem_morethanprice = client.query('''
        SELECT
            c.id as claimID,
            p.extOrderId as orderId,
            cast(json_extract(i.details,'$.Raw.priceUSD') as int64)/100 / (i.quantity*i.price) * claimAmount as claimAmountUSD,
            cast(json_extract(i.details,'$.Raw.priceUSD') as int64)/100 as priceUSD,
            c.*
        FROM `axinan-data.shopee_XB_Claim.claim` as c
        left join
        `axinan-data.shopee_XB_Claim.item` as i
        on c.item_id = i.itemId and i.policyId = c.policyId
        left join
        `axinan-data.shopee_XB_Claim.Policy` as p
        on i.policyId = p.id
        left join
        `axinan-data.shopee_XB_Claim.Order` as o
        on p.extOrderId = o.extOrderId

        where cast(json_extract(i.details,'$.Raw.priceUSD') as int64)/100 / (i.quantity*i.price) * claimAmount > cast(json_extract(i.details,'$.Raw.priceUSD') as int64)/100 + 1
        and i.deletedAt is null and p.deletedAt is null and o.deletedAt is null
        and total_items = 1
        order by orderId
        ''').to_dataframe()

    if df_singleitem_morethanprice.empty:
        print("Pass: single item claim amount check")
    else:
        df_singleitem_morethanprice.to_csv("Data_singleitem_morethanprice.csv")
        print("To check single item claim amount > price more than $1 USD: Data_singleitem_morethanprice.csv")

    # for claims with multiple items, check claim amount is no more than total value of the claimed items
    df_multiitem_morethanprice = client.query('''
            SELECT
                c.id as claimID,
                p.extOrderId as orderId,
                claimAmount as claimAmountLocal,
                max_claim.sum_item_claim as total_item_value,
                c.*
            FROM `axinan-data.shopee_XB_Claim.claim` as c
            left join
            `axinan-data.shopee_XB_Claim.item` as i
            on c.item_id = i.itemId and i.policyId = c.policyId
            left join
            `axinan-data.shopee_XB_Claim.Policy` as p
            on i.policyId = p.id
            left join
            `axinan-data.shopee_XB_Claim.Order` as o
            on p.extOrderId = o.extOrderId
            left join
            (
            SELECT
            id,
            sum(i.quantity*i.price) as sum_item_claim
            FROM `axinan-data.shopee_XB_Claim.claim_expanditem` as c
            left join
            `axinan-data.shopee_XB_Claim.item` as i
            on c.item_id = i.itemId and i.policyId = c.policyId
            group by id
            ) as max_claim
            on c.id = max_claim.id

            where  i.deletedAt is null and p.deletedAt is null and o.deletedAt is null
            and total_items > 1 and claimAmount>max_claim.sum_item_claim
            order by orderId
            ''').to_dataframe()

    if df_multiitem_morethanprice.empty:
        print("Pass: multi item claim amount check")
    else:
        df_multiitem_morethanprice.to_csv("Data_multiitem_morethanprice.csv")
        print("To check multi item claim, the claim amount is more than value of all items: Data_multiitem_morethanprice.csv")

    # for refund succeed, check there is a refund amount,
    # ignore 0 claim amounts, e.g. 0 after excess
    df_succeed_withclaimAmount_norefundAmount = client.query('''
            SELECT
            *
            FROM `axinan-data.shopee_XB_Claim.claim`
            where refundStatus = "succeed" and claimAmount	>0 and refundAmount = 0
            and Proof_type not like '%50%'
            ''').to_dataframe()

    if df_succeed_withclaimAmount_norefundAmount.empty:
        print("Pass: 0 refund amount check")
    else:
        df_succeed_withclaimAmount_norefundAmount.to_csv("Tech_succeed_withclaimAmount_norefundAmount.csv")
        print(
            "To check claims with refundStatus succeed, none-0 claimAmount and 0 refundAmount and not due to Excess: Tech_succeed_withclaimAmount_norefundAmount.csv")

    # for claims without excess, e.g. reasons without 50+0 usd
    # verify refund amount is 50%
    df_refund50 = client.query('''
            SELECT
            *
            FROM `axinan-data.shopee_XB_Claim.claim`
            where refundStatus = "succeed" and claimAmount	>0
            and abs(claimAmount/refundAmount - 2) > 0.001
            and refundAmount != 0
            and Proof_type not like '%50%' #excess checked separately
            and Proof_type != 'exclude_insurance' #checked separately
            ''').to_dataframe()

    if df_refund50.empty:
        print("Pass: refund amount 50% check")
    else:
        df_refund50.to_csv("refund50.csv")
        print(
            "To check refund amount != 50% of claim amount: Tech_refund50.csv")

    # verify invoice amount is consistant with claim amount
    df_invoiceV2 = client.query('''
                SELECT
                c.id as claimId,
                i.amount as invoiceV2_amount,
                c.refundamount as claim_refundamount

                FROM `axinan-data.shopee_XB_Claim.InvoiceV2` as i
                left join
                `axinan-data.shopee_XB_Claim.claim` as c
                on i.claimId = c.id

                where type = 'claim'
                and abs(i.amount - c.refundamount) > 100
                ''').to_dataframe()

    if df_invoiceV2.empty:
        print("Pass: invoiceV2 table check")
    else:
        df_invoiceV2.to_csv("invoiceV2.csv")
        print(
            "To check invoiceV2 amount different from refundamount: df_invoiceV2.csv")

    # verify invoice batch amount is consistent with claim amount
    df_invoiceBatchV2 = client.query('''
                SELECT
                *
                from
                (
                  SELECT
                  ib.id,
                  cast(json_query(ib.details, '$.Raw.totalRefundAmountUSD') as int64) as invoice_refund_USD,
                  sum(cast(c.refundAmountUSD as int64)) as claim_refundamount

                  FROM `axinan-data.shopee_XB_Claim.InvoiceBatchV2` as ib
                  left join
                  `axinan-data.shopee_XB_Claim.InvoiceV2` as i
                  on ib.id = i.batchid
                  left join
                  `axinan-data.shopee_XB_Claim.claim` as c
                  on i.claimId = c.id

                  where i.type = 'claim' and ib.type = 'claim' and c.refundstatus = 'succeed'
                  group by ib.id, invoice_refund_USD
                )
                where invoice_refund_USD != claim_refundamount
                                    ''').to_dataframe()

    if df_invoiceBatchV2.empty:
        print("Pass: invoiceBatchV2 table check")
    else:
        df_invoiceBatchV2.to_csv("invoiceBatchV2.csv")
        print("To check invoiceBatch V2 refund amount != sum refundstatus = 'succeed' claim refund amount : df_invoiceV2.csv")

    # check if invoices are generated within 7 days of approval
    df_invoicenotgenerated7days = client.query('''
                SELECT
                TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), c.axinanApprovedAt, DAY) as invoice_pending_days,
                c.*
                FROM `axinan-data.shopee_XB_Claim.claim` as c
                left join
                `axinan-data.shopee_XB_Claim.InvoiceV2` as i
                on i.claimId = c.id
                where c.status = 'approved' and i.createdAt is null
                and TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), c.axinanApprovedAt, DAY) > 7
                                    ''').to_dataframe()

    if df_invoicenotgenerated7days.empty:
        print("Pass: generate invoice in 7 days check")
    else:
        df_invoicenotgenerated7days.to_csv("Tech_invoicenotgenerated7days.csv")
        print("To check invoice not generated in 7 days : Tech_invoicenotgenerated7days.csv")

    #verify 7 days  no reason return is only for TW
    df_7daysTWD = client.query('''
        SELECT *
        FROM `axinan-data.shopee_XB_Claim.claim`
        where Proof_type = '7 days no reason to return' and currency != 'TWD'
                                        ''').to_dataframe()

    if df_7daysTWD.empty:
        print("Pass: 7 days no reason to return only in TW")
    else:
        df_7daysTWD.to_csv("Data_7daysnoreasonTW.csv")
        print("To 7 days no reason not only in TW : Data_7daysnoreasonTW.csv")



## uncomment and run the required part

# update_item_policy_order()
#
# update_shop()
#
# update_claim()
#
# update_invoice()

check_claim()





print("--- %s seconds ---" % (time.time() - start_time))








