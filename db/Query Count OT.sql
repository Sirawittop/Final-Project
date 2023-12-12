-- Active: 1701774805198@@127.0.0.1@8889@NursePlan
-- Active: 1701774805198@@127.0.0.1@8889@NursePlan
SELECT nurseid, 
        SUM(CASE WHEN Day1 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day2 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day3 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day4 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day5 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day6 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day7 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day8 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day9 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day10 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day11 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day12 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day13 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day14 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day15 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day16 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day17 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day18 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day19 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day20 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day21 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day22 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day23 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day24 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day25 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day26 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day27 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day28 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day29 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day30 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END +
            CASE WHEN Day31 IN (10 , 11 , 12 , 13 , 14 ,15 ,16 ,17 ,18 ,20, 21, 22) THEN 1 ELSE 0 END) AS 'Day Shifts'
FROM plans
GROUP BY nurseid;
