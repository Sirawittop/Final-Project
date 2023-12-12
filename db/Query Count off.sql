-- Active: 1701774805198@@127.0.0.1@8889@NursePlan
SELECT nurseid, 
        SUM(CASE WHEN Day1 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day2 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day3 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day4 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day5 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day6 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day7 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day8 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day9 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day10 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day11 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day12 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day13 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day14 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day15 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day16 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day17 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day18 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day19 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day20 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day21 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day22 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day23 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day24 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day25 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day26 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day27 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day28 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day29 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day30 = 4 THEN 1 ELSE 0 END +
            CASE WHEN Day31 = 4 THEN 1 ELSE 0 END) AS 'Day Shifts'
FROM plans
GROUP BY nurseid;
