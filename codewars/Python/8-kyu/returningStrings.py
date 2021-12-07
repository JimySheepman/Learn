def greet(name):
    main_string = "Hello, <name> how are you doing today?"
    parse_string = main_string.split("<name>")
    parse_string = name.join(parse_string)
    return parse_string