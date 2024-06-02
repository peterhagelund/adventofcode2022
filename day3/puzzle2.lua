local function main()
    local f = io.open("puzzle_input.txt", "r")
    if not f then
        os.exit(1)
    end
    local sum = 0
    local group = {"", "", ""}
    local index = 1
    for line in f:lines() do
        group[index] = line
        index = index + 1
        if index == 4 then
            for i = 1, #group[1] do
                local c = string.sub(group[1], i, i)
                local j, _ = string.find(group[2], c)
                local k, _ = string.find(group[3], c)
                if j and k then
                    if c >= "a" and c <= "z" then
                        sum = sum + string.byte(c) - string.byte("a") + 1
                    else
                        sum = sum + string.byte(c) - string.byte("A") + 27
                    end
                    break
                end
            end
            index = 1
        end
    end
    f:close()
    print(string.format("sum = %d", sum))
end

main()
