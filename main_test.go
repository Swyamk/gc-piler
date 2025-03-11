package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const testcaseDirectory = "testcases/"

var testPrograms = []string{
	//"multiline_map",
	"multiline_string",
	"var_string",
	"var_multi",
	"string_args",
	"defer",
	"iota",
	"map_struct",
	"for_range_map_key_value",
	"for_range_map_value",
	"for_range_map_key",
	"string_map",
	"struct",
	"var2",
	"for_range_both",
	"for_endless",
	"for_range_single",
	"for_range_list",
	"for_regular",
	"trimspace",
	"printf",
	"multiple",
	"prefix",
	"contains",
	"if",
	"var",
	"output_char2",
	"var_zero",
	"output_char",
	"const",
	"continue",
	"types",
	"if_minus_one",
	"goto",
	"switch",
	"hello",
	"sprintf",
	"functions_test", 
    "maps_test",      
    "slices_test",
	"struct_methods_test", 
    "geometry_test",       
    "concurrency_test",
}

