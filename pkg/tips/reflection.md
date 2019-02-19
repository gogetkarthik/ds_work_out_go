# Go Personal Note

### Reflection
> reflect.TypeOf accepts **interface{}** returns **Dynamic Type**
    t := reflect.TypeOf(3) 
    
> reflect.ValueOf accepts **interface{}** returns **Dynamic value**
    <p> t := reflect.ValueOf(3)
    <p> t.String() //if the value is not string it will give the type