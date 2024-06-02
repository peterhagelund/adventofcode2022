local function main()
    local f = io.open("puzzle_input.txt", "r")
    if not f then
        os.exit(1)
    end
    local sum = 0
    for line in f:lines() do
        local l = #line / 2
        local c1 = string.sub(line, 1, l)
        local c2 = string.sub(line, l + 1, #line)
        for i = 1, #c1 do
            local c = string.sub(c1, i, i)
            local j, _ = string.find(c2, c)
            if j then
                if c >= "a" and c <= "z" then
                    sum = sum + string.byte(c) - string.byte("a") + 1
                else
                    sum = sum + string.byte(c) - string.byte("A") + 27
                end
                break
            end
        end
    end
    f:close()
    print(string.format("sum = %d", sum))
end

main()
