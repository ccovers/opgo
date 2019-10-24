package main

import (
	"fmt"
)

/*
* 朴素贝叶斯分类算法
* 什么是分类：
	分类是一种重要的数据分析形式，它提取刻画重要数据类的模型。这种模型称为分类器，预测分类的（离散的，无序的）类标号。例如医生对病人进行诊断是一个典型的分类过程，医生不是一眼就看出病人得了哪种病，而是要根据病人的症状和化验单结果诊断病人得了哪种病，采用哪种治疗方案。再比如，零售业中的销售经理需要分析客户数据，以便帮助他猜测具有某些特征的客户会购买某种商品。
* 如何进行分类：
	数据分类是一个两阶段过程，包括学习阶段（构建分类模型）和分类阶段（使用模型预测给定数据的类标号）
* 贝叶斯分类的基本概念：
	贝叶斯分类法是统计学分类方法，它可以预测类隶属关系的概率，如一个给定元组属于一个特定类的概率。贝叶斯分类基于贝叶斯定理。朴素贝叶斯分类法假定一个属性值在给定类上的概率独立于其他属性的值，这一假定称为类条件独立性。
* 贝叶斯定理：
	贝叶斯定理特别好用，但并不复杂，它解决了生活中经常碰到的问题：已知某条件下的概率，如何得到两条件交换后的概率，也就是在已知P(A|B)的情况下如何求得P(B|A)的概率。P(A|B)是后验概率（posterior probability），也就是我们常说的条件概率，即在条件B下，事件A发生的概率。相反P(A)或P(B)称为先验概率（prior probability·）。贝叶斯定理之所以有用，是因为我们在生活中经常遇到这种情况：我们可以很容易直接得出P(A|B)，P(B|A)则很难直接得出，但我们更关心P(B|A)，贝叶斯定理就为我们打通从P(A|B)获得P(B|A)的道路。

* 朴素贝叶斯分类的思想和工作过程：
	朴素贝叶斯分类的思想真的很朴素，它的思想基础是这样的：对于给出的待分类项，求解此项出现的条件下各个类别出现的概率，哪个最大，就认为此待分类属于哪个类别。
*/

func main() {

}
