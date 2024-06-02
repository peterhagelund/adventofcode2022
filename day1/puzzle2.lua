local function main()
    local f = io.open("puzzle_input.txt", "r")
    if not f then
        os.exit(1)
    end
    local allCalories = {}
    local elfCalories = 0
    for line in f:lines() do
        if #line > 0 then
            elfCalories = elfCalories + tonumber(line)
        else
            table.insert(allCalories, elfCalories)
            elfCalories = 0
        end
    end
    if elfCalories > 0 then
        table.insert(allCalories, elfCalories)
    end
    f:close()
    table.sort(allCalories, function(a, b) return a > b end)
    local topThreeCalories = allCalories[1] + allCalories[2] + allCalories[3]
    print(string.format("top three calories = %d", topThreeCalories))
end

main()
