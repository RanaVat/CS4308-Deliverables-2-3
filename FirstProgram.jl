# This program showcases Multiple Dispatch in Julia
# Based on the following youtube video: https://www.youtube.com/watch?v=hesjQz__yb8

# overloaded function give_type
function give_type(x::String)
    return "You gave me a String!"
end 

function give_type(x::Int64)
    return "You gave me an Int64!"
end

function give_type(x::Float16)
    return "You gave me a Float16!"
end

function give_type(x::Any)
    return "You gave me a miscellaneous type!"
end

# Multiple Dispatch

function truly_generic(x)
    println("$x: ", give_type(x))
end

truly_generic(64)
truly_generic("Hello World!")
truly_generic(16.0)
truly_generic(true)