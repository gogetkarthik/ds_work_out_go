package tips

//Use a pointer if you need to pass something to be modified.
//Use a pointer if you need to determine if something was unset/nil.
//Use a pointer if you are using a type that has methods with pointer receivers.

//Returning structs vs. pointers seems to be roughly the same speed, but passing pointers to functions down the lines is siginificantly faster. Although not on a level it would matter

//Slices, maps, channels, strings, function values, and interface values are implemented with pointers internally, and a pointer to them is often redundant.

//Value types
//int, float, string, bool, structs
// Reference Tyoe
//slices, maps, channels, pointers, functions
