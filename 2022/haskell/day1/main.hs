splitString :: Char -> [Char] -> [[Char]]
splitString _ [] = []
splitString sep str = 
    let (left, right) = break (==sep) str 
    in left : splitString sep (drop 1 right)


splitList:: a -> [a] -> [[a]]
splitList _ [] = []
splitList sep str = 
    let (left, right) = break (EQ sep) str 
    in left : splitList sep (drop 1 right)




main :: IO ()
main = do
    input <- readFile "input-sample.txt"
