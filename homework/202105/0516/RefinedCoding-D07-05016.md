# RefinedCoding-D07-05016      
- https://talks.golang.org/2012/splash.article      
- 谷歌的Go语言-论软件工程服务中的语言设计      
      
## 1.概要      
      
## 2.简介      
      
## 3.Go at Google      
      
## 4.痛点      
- 编译慢      
- 无法控制编译依赖      
- 程序员用语言的不同子集      
- 代码晦涩难懂      
- 重复劳动      
- 升级代价大      
- 版本偏离      
- 自动化工具难写      
- 交叉编译      
      
## 5.C/C++依赖      
      
## 6.Go时代      
      
## 7.Go依赖      
      
## 8.包管理      
      
## 9.远程包      
      
## 10.Syntax语法      
- 用var开头定义变量的理由竟然是容易解析      
- 解析int x = 10;在编译原理里面不是很简单么      
- var x int = 10，多一个var，感觉累赘      
      
## 11.变量命名      
- 看了去语言的设计，终于明白我们用到的方法都是大写开头了      
- 方法是有主体的函数？      
      
## 12.Semantics语义      
      
## 13.Concurrency并发      
      
## 14.Garbage Collection      
      
## 15.组合优于继承      
      
## 16.错误      
      
## 17.工具      
      
## 18.结论      
      
## 19.总结      
      
## Teminology      
conceived,构想        
landscap,景观        
addressed,已解决        
comprise,包括        
complilation,顺应        
aspect,方面        
rigorous,严格的        
compiled,已编译        
unimaginative,缺乏想象力        
contradictory,矛盾的        
breakthrough,突破        
nontheless,尽管如此        
comprising,包括        
clumsy,笨拙        
clumsiness,笨拙        
thereby,从而        
rephrase,改写        
methologies,方法论        
derigueur,死神        
depedencies,礼节        
skew,歪斜        
indentation,缩进        
snippet,片段        
invocation,调用        
heterogeneous,异质        
substaintial,实质性的        
ubiquitous,无处不在        
discipline,纪律        
dramatically,显着地        
exacerbates,加剧        
granularity,粒度        
intricate,错综复杂        
instrumented,仪器化的        
concatenated,级联的        
shrank,缩水        
precursor,前体        
conveived,对流        
procedural,程序        
particularly,特别        
radical,激进的        
advent,来临        
perspective,看法        
syntactically,句法上        
semanticaly,语义上        
comprising,包括        
unused,没用过        
guarantee,保证        
precise,精确的        
extraneous,外来的        
transitively,传递地        
motivating,激励        
clause,条款        
reminiscent,使人想起        
opposed,反对的        
export,出口        
verbose,冗长的        
interacive,互动的        
cycles,周期数        
entangling,纠缠        
swaths,地带        
subpieces,子件        
bloating,腹胀        
occasional annoyance,偶尔的烦恼        
boundaries,界线        
satisfactorily,满意地        
hygiene trumps,卫生王牌        
conversion,转换        
honors,荣誉        
concise,简洁的        
clarity,明晰        
renamed,重命名        
collisions,碰撞        
arbitrary,随意的        
decentralized,去中心化        
syntax,句法        
semantics,语义学        
arguably,可以说        
clarity,明晰        
quirks,怪癖        
symbol,象征        
grammar,语法        
declaration,宣言        
introduced,引进        
eliminates ambiguity,消除歧义        
straightforward,直截了当        
receiver,接收者        
first-class,头等舱        
default,默认        
deliberate,商榷        
flaws,缺陷        
disentangle,解开        
leads to,导致        
aspect,方面        
mitigating,减轻        
variadic,可变的        
initial,最初的        
upper,上        
types,类型        
easy,简单的        
burdensome,繁重的        
clarityuniverse predeclared,明晰宇宙已预先声明        
scope,范围        
lexically,词汇上        
qualifier,限定词        
unambifuous,明确的        
decouple,解耦        
signature,签名        
invocation,调用        
semantics,语义学        
proceduralpointers,程序指针        
accustomed,习惯的        
audience,观众        
rooting,生根        
arithmetic,算术        
bounds,界线        
aliases,别名        
linguistic,语言学        
perspective,看法        
embodies,体现        
variant,变体        
CSP,CSP        
familiarity,熟识        
predecessor,前任        
orthogonal,正交的        
ordinary,普通的        
composition,作品        
cryptographic,密码学        
canonical,典范        
caveat,警告        
presence,在场        
idiomatic,惯用的        
write-once,一次写入        
familiarity,熟识        
forbid,禁止        
convention,习俗        
motto,座右铭        
controversial,有争议的        
profound,深刻的        
conversely,反过来        
resource management,资源管理        
overhead,高架        
nonetheless,尽管如此        
outweight,超重        
borne,承担        
overheads,间接费用        
tuning,调音        
mitigate,缓解        
layout,布局        
second alloaction,第二次同种异体        
indirection,间接的        
second-order,二阶        
arena,竞技场        
prearrangement,预先安排        
dynamic,动态的        
interior,内部的        
idiomatic,惯用的        
colletor,同事        
interiorarena,室内竞技场        
overheadmandate,管理费用        
wisely,明智地        
mthods,方法        
hierarchy,等级制度        
intentional,故意的        
oveused,隐瞒        
interfaces,介面        
despite,尽管      
fluidity,流动性      
subclassing,子类化      
identical,完全相同的      
uniformity,均匀度      
orthogonally,正交地      
extreme,极端      
Plan 9,方案9      
uniformity,均匀度      
abound,盛产      
wordaday,世代相传      
argue,争论      
neglected,被忽视      
brittle,脆      
overdesign,过度设计      
upside,上行空间      
dawn,黎明      
composition,作品      
trivial,不重要的      
comprehensivle boundaries,理解边界      
compelementary,互补的      
chaining,链式      
composition,作品      
accustomed,习惯的      
adaptability,适应性      
dependency,依赖      
organically,有机地      
predetermined,预定的      
distrubs,杂草      
functions,职能      
widespread,广泛      
looser,宽松的      
decoupled,解耦的      
exception,例外      
convential,传统的      
clumsy,笨拙      
alongside,并排      
deliberatecritics,蓄意批评      
exceptional,非同凡响      
linguistic,语言学      
distorts,扭曲      
interlaces,隔行扫描      
contrast,对比      
verbose,冗长的      
explict error checking,显式错误检查      
myriad,无数      
lexer,词法分析器      
manipulate,操纵      
consequences,结果      
eliminating,消除      
lay,躺下      
presubmit,预先提交      
foresee,预见      
refactoring,重构      
semantically,语义上      
rewrite,改写      
right-hand side,右手边      
slice expression,切片表达      
cononical,圆锥形的      
intricate,错综复杂      
semicolons,分号      
sweeping,扫地      
entries,参赛作品      
radically,根本      
rolled out,推出      
deprecated,不推荐使用      
infeasible,不可行的      
extractor,提取器      
compatibility,兼容性      
rarely,很少      
ecosystem,生态系统      
delivers,交付      
forms,形式      
opportunity,机会      
significant,重大      
      
      
      
      
      
      
      
      
      
      
      
      
      
      