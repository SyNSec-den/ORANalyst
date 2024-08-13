file(REMOVE_RECURSE
  "libe2sim.a"
  "libe2sim.pdb"
)

# Per-language clean rules from dependency scanning.
foreach(lang C CXX)
  include(CMakeFiles/e2sim_static.dir/cmake_clean_${lang}.cmake OPTIONAL)
endforeach()
