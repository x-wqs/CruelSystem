6: Large Scaling system uses MapReduce for its high performance, strong fault telorance, simple interface, and time of one version iteration decreases from months to days.

8: MapReduce has been successfully used in Google for many purpose, for three reasons: easy to use, a large amounts of problems easily attributed to this model, scaling large clusters composed of thousands of machines.

In practice, network bandwidth is scarce, so numbers of optimization about reducing data sent cross network is needed, which we can see in locality optimization.
And little redundant execution can speed up the whole process apparently, such as backup tasks, which tackles the straggler problem caused by slow machines.

Analysis of Appendix A:
```cpp
#include "mapreduce/mapreduce.h"

// mapper class
class MyMapper : public Mapper {
 public:
  virtual void Map(const MapInput& input) {
    auto key = Deal(input);
    Emit(key, 1);
  }
}
REGISTER_MAPPER(MyMapper)

// reducer class
class MyReducer : public Reducer {
 public:
  Virtual void Reduce(ReduceInput** input) {
    auto sum_value;
    for (auto value : input->values()) sum_value += value;
      Emit(sum_value);
    }
}
REGISTER__REDUCER(MyReducer)

int main(int argc, char** argv) {
  
  MapReduceSpecification spec;
  // set mapper
  for (int i = 1; i < argc; i++) {
    MapReduceInput* input = spec.add_input();
    input->set_format("text");
    input->set_filepattern(argv[i]);
    input->set_mapper_class("MyMapper");
  }

  // set reducer
  MapReduceOutput* out = spec.output();
  out->set_filebase("/outprefix");
  out->set_num_tasks(100);
  out->set_format("text");
  out->set_reducer_class("Reducer");

  out->set_combiner_class("Reducer");

  spec.set_machines(2000);
  spec.set_map_megabytes(100);
  spec.set_reduce_megabytes(100);

  MapReduceResult result;
  if (!MapReduce(spec, &result)) abort();

  return 0;
}
```

reference: https://www.youtube.com/watch?v=Rz8JCS9TfOQ
