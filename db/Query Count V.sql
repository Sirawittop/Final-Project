-- Active: 1701774805198@@127.0.0.1@8889@NursePlan
-- Active: 1701774805198@@127.0.0.1@8889@NursePlan
SELECT nurseid, 
        SUM(CASE WHEN Day1 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day2 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day3 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day4 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day5 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day6 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day7 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day8 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day9 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day10 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day11 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day12 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day13 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day14 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day15 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day16 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day17 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day18 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day19 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day20 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day21 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day22 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day23 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day24 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day25 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day26 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day27 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day28 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day29 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day30 = 5 THEN 1 ELSE 0 END +
            CASE WHEN Day31 = 5 THEN 1 ELSE 0 END) AS 'Day Shifts'
FROM plans
GROUP BY nurseid;
