-- Active: 1701774805198@@127.0.0.1@8889@NursePlan
-- Active: 1701774805198@@127.0.0.1@8889@NursePlan
SELECT nurseid, 
        SUM(CASE WHEN Day1 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day2 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day3 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day4 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day5 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day6 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day7 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day8 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day9 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day10 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day11 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day12 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day13 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day14 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day15 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day16 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day17 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day18 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day19 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day20 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day21 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day22 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day23 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day24 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day25 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day26 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day27 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day28 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day29 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day30 = 6 THEN 1 ELSE 0 END +
            CASE WHEN Day31 = 6 THEN 1 ELSE 0 END) AS 'Day Shifts'
FROM plans
GROUP BY nurseid;
