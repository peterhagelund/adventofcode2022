local function main()
    local f = io.open("puzzle_input.txt", "r")
    if not f then
        os.exit(1)
    end
    local maxCalories = 0
    local elfCalories = 0
    for line in f:lines() do
        if #line > 0 then
            elfCalories = elfCalories + tonumber(line)
        else
            maxCalories = math.max(maxCalories, elfCalories)
            elfCalories = 0
        end
    end
    if elfCalories > 0 then
        maxCalories = math.max(maxCalories, elfCalories)
    end
    f:close()
    print(string.format("max calories = %d", maxCalories))
end

main()
