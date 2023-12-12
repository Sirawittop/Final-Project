SELECT nurseid, 
        SUM(CASE WHEN Day1 IN (2, 10 , 13  , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day2 IN (2, 10 , 13  , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day3 IN (2, 10 , 13  , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day4 IN (2, 10 , 13  , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day5 IN (2, 10 , 13  , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day6 IN (2, 10 , 13  , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day7 IN (2, 10 , 13  , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day8 IN (2, 10 , 13  , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day9 IN (2, 10 , 13  , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day10 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day11 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day12 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day13 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day14 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day15 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day16 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day17 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day18 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day19 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day20 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day21 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day22 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day23 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day24 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day25 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day26 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day27 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day28 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day29 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day30 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END +
            CASE WHEN Day31 IN (2 , 10 , 13 , 20) THEN 1 ELSE 0 END) AS 'Day Shifts'

        

FROM plans
GROUP BY nurseid;
