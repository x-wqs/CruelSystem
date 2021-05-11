Map: divide in degree of inputs, generate intermediates: {key, value}
Reduce: divide in degree of intermediates, deal with {key, value}s with same key
(pairs with same key reduce to one pair, that's reduce means)

map(struct input):
	generate({input.GetKey(), input.GetValue()})

reduce(string key, vector<T> value_vec):
	T value_sum;
	for (auto value : value_vec) value_sum.add(value);
	generate({key, value_sum})

2.3: common examples: word count, distribute grep, distribute sort, inverted index

3: usage of system: users submit jobs to scheduling system(one job consists of multiple tasks)
3.1: M mapper + R reducer + 1 master
		 usually input_size / M = 64M
3.2: master mantain locations of each split of immediates, a table of M * R stats
3.3: worker failure: re-execute when map done, map in progress, reducer in progress(reducer done don't need to re-execute)
     master failure: shutdown

