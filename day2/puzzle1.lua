local rock, paper, scissors = 1, 2, 3

local function determineOutcome(opponent, me)
    if opponent == rock then
        if me == rock then
            return 1 + 3
        elseif me == paper then
            return 2 + 6
        else
            return 3 + 0
        end
    elseif opponent == paper then
        if me == rock then
            return 1 + 0
        elseif me == paper then
            return 2 + 3
        else
            return 3 + 6
        end
    else
        if me == rock then
            return 1 + 6
        elseif me == paper then
            return 2 + 0
        else
            return 3 + 3
        end
    end
end

local function main()
    local f = io.open("puzzle_input.txt", "r")
    if not f then
        os.exit(1)
    end
    local hands = {
        A = rock,
        B = paper,
        C = scissors,
        X = rock,
        Y = paper,
        Z = scissors,
    }
    local points = 0
    for line in f:lines() do
        local opponent = hands[string.sub(line, 1, 1)]
        local me = hands[string.sub(line, 3, 3)]
        points = points + determineOutcome(opponent, me)
    end
    f:close()
    print(string.format("points = %d", points))
end

main()
