-- Active: 1701774805198@@127.0.0.1@8889@NursePlan
SELECT nurseid, 
        SUM(CASE WHEN Day1 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day2 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day3 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day4 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day5 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day6 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day7 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day8 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day9 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day10 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day11 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day12 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day13 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day14 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day15 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day16 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day17 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day18 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day19 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day20 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day21 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day22 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day23 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day24 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day25 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day26 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day27 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day28 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day29 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day30 = 7 THEN 1 ELSE 0 END +
            CASE WHEN Day31 = 7 THEN 1 ELSE 0 END) AS 'Day Shifts'
FROM plans
GROUP BY nurseid;
