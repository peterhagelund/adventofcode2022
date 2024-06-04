local function main()
    local f = io.open("puzzle_input.txt", "r")
    if not f then
        os.exit(1)
    end
    local sum = 0
    for line in f:lines() do
        local it = string.gmatch(line, "(%d+)")
        local s1, e1, s2, e2 = tonumber(it()), tonumber(it()), tonumber(it()), tonumber(it())
        if (s1 >= s2 and s1 <= e2) or (e1 >= s2 and e1 <= e2) or (s2 >= s1 and s2 <= e1) or (e2 >= s1 and e2 <= e1) then
            sum = sum + 1
        end
    end
    f:close()
    print(string.format("sum = %d", sum))
end

main()
