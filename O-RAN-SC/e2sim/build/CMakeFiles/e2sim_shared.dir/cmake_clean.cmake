file(REMOVE_RECURSE
  "libe2sim_shared.pdb"
  "libe2sim_shared.so"
)

# Per-language clean rules from dependency scanning.
foreach(lang C CXX)
  include(CMakeFiles/e2sim_shared.dir/cmake_clean_${lang}.cmake OPTIONAL)
endforeach()
