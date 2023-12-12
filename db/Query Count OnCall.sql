-- Active: 1701774805198@@127.0.0.1@8889@NursePlan
SELECT nurseid, 
        SUM(CASE WHEN Day1 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day2 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day3 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day4 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day5 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day6 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day7 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day8 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day9 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day10 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day11 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day12 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day13 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day14 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day15 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day16 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day17 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day18 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day19 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day20 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day21 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day22 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day23 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day24 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day25 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day26 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day27 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day28 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day29 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day30 = 8 THEN 1 ELSE 0 END +
            CASE WHEN Day31 = 8 THEN 1 ELSE 0 END) AS 'Day Shifts'
FROM plans
GROUP BY nurseid;
