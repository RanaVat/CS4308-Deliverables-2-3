# This program showcases dictionaries in Julia

Storage = Dict(
    "Unit A"=>"Baseball Bats",
    "Unit B"=>"Old Furniture",
    "Unit C"=>"Anime Figurines",
    "Unit D"=>"Computer Parts"
)

# Prints the value of the passed key
Storage["Unit D"]

# Prints all keys and values in the dictionary
for item in Storage
    print(item.first)
    print(" : ")
    println(item.second)
end

# Changes value using key
Storage["Unit D"] = "Modern Art Collection"

# Adds a new pair
Storage["Unit Z"] = "Classified Government Documents"

# Removes a pair
pop!(Storage, "Unit B")