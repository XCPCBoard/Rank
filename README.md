# rank
## 1. 基础分

$BaseRating =ProblemScorce*0.5+ratingScorce*0.4+BlogScore*0.1$

$ProblemScorce = PassPloblemSum$

$ratingScorce=(AtcodeRating+CodeforcesRating)*0.1$

$BlogScorce=BlogNum*2$

## 2. Rating 计算

$Rating$ 将以(天\周\月)为单位作为一场比赛，进行迭代计算。

### 2.1 模块分数
#### 2.1.1 ProblemScore
$ProblemScore=easy∗1+basic∗2+advanced∗3+hard∗4+unknown∗2$
#### 2.1.2 ratingScore
$codeforces ：$

$d=rating_{cf_{new}}-rating_{cf}$

$$
ratingScore_{cf}=
\begin{cases}
1,\quad rating_{cf}\leq 600,d \geq 1\\
\frac{rating_{cf}}{400}+\frac{rating_{cf}*d}{20000}, \quad rating_{cf}\geq 601,d \geq 1\\
0,\quad d \leq 0
\end{cases}
$$

$Atcoder:$

$d=rating_{atc_{new}}-rating_{atc}$

$$
ratingScore_{atc}=
\begin{cases}
1,\quad rating_{atc}\leq 400,d \geq 1\\
\frac{rating_{atc}}{400}\cdot(1+\frac{d}{50}), \quad rating_{atc}\geq 401,d \geq 1\\
\frac{rating_{atc}}{400}\cdot(1+\frac{d}{20}), \quad rating_{atc}\geq 1000,d \geq 1\\
0,\quad d \leq 0
\end{cases}
$$

#### 2.1.3 BlogScore
$BlogScore=\frac{\sum^{n}_{i=1} BlogScore_i}{n}$

#### 2.1.4 AttendanceScore
$AttendanceScore=单位周期内出勤分钟数$


### 2.2 预期胜率
ELO积分预期胜率计算公式

$P(D)=\frac{1}{2}+\int_0^D \frac{1}{\delta\sqrt{2\pi}} \cdot	 e^{\frac{-x^2}{2\delta ^2}}dx$

利用最小二乘法得到实际应用公式，其中 $D$ 代表分差。

$P(D)=\frac{1}{1+10^{\frac{D}{400}}}$

### 2.3 周期表现 rating 计算
#### 2.3.1 考虑 1 V 1
$R_A:player A 的 rating ~~~~~~~~~~~~~~~~~~~~~~~~~~~R_B:player B 的 rating$

$E_A=\frac{1}{1+10^{\frac{R_B-R_A}{400}}}~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~E_B=\frac{1}{1+10^{\frac{R_A-R_B}{400}}}$

$E_A+E_B=1$

$S_A=\frac{problemScore_A}{problemScore_A+problemScore_B} * 0.4+\frac{ratingScore_A}{ratingScore_A+ratingScore_B}*0.3+\frac{blogScore_A}{blogScore_A+blogScore_B}*0.2+\frac{AttendanceScore_A}{AttendanceScore_A+AttendanceScore_B}*0.1$

$S_B=\frac{problemScore_B}{problemScore_A+problemScore_B} * 0.4+\frac{ratingScore_B}{ratingScore_A+ratingScore_B}*0.3+\frac{blogScore_B}{blogScore_A+blogScore_B}*0.2+\frac{AttendanceScore_B}{AttendanceScore_A+AttendanceScore_B}*0.1$

$S_A+S_B=1$

$R_{A_{new}}=R_A+K \cdot (S_A -E_A)$

$K$ 暂定为 $32$，实际上 $K$ 将随着用户 $Rating$ 的增加而减小。


#### 2.3.2 考虑 1 V n
$R_{A_{new}}=R_A+K \cdot P_A$



$P_A=\sqrt[x]{\prod^{x}_{R_i<R_A}(S_{Ai}-E_{Ai})}-\sqrt[y]{\prod^{y}_{R_i>R_A}(E_{Ai}-S_{Ai})}$



由于$(S_{A_i}-E_{A_i})$并不是全为正数，因此通过分别计算对应的值做差为 $P_A$。

#### 2.3.3  rating 修正调整
第一次：

$adjust=\frac{-1-\sum K_i*P_i}{n}$

$R_i=R_i+adjust$

保证所有人的平均变化接近 0 并且在 0 以下。

第二次：

$m=min(n,4\sqrt{n})$


$adjust=min(max(\frac{-1-\sum K_i*P_i}{m},-10),0)$

取一个合理的 $adjust$ 使得前 $m$ 个人的平均变化为 0。


$R_i=R_i+adjust~~~~(i \leq m)$

## 3. 后记
本文参考了codeforce，atcoder，Elo rating system$ 的 $rating​ 规则，以周期统计数据替代比赛场景，并根据应用场景进行修改，目前未进行样本测试。

可能测试后，还会对参数以及公式大改QAQ

> 参考
> https://en.wikipedia.org/wiki/Elo_rating_system
>
> https://www.luogu.com.cn/blog/ak-ioi/cf-at-rating