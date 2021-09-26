阅读区+理解区
不在于回答得快，而在于要深思熟虑理解问题之后找到一个方案
从对面的角度看，想考察什么？不仅是系统设计的能力
问问题的能力，合作能力，压力下的反应，从模糊走到清晰的能力，会不会over engineering导致high cost

四步走：
1 理解问题（by asking questions)，找到scope
	题目：做一个系统；行动：列出可能的scenario和feature；问题：做哪些部分
	问题：有多少用户
	问题：期望scale到多少用户，3/6/12月
	行动：列出可能的现有的tech stack，现有的service 问题：我们可以assume哪些已经有且可以加以利用
2 high level设计让别人buy in（框架图）
	initial blueprint讲给面试官听
	画盒子（api，web server，data store）
	问题：back of the envelope estimation是否有必要；（如有必要）看是否能满足1里面的scaling要求
	过几个用户案例，帮助设计并验证设计
	用问题大小判断是否要设计到api level还是database schema level(大：“Design Google search engine”；小：a multi-player poker game,)，问面试官来确定
