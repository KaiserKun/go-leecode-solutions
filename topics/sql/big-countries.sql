SELECT name, population, area
FROM World as w
where w.area >= 3000000 OR w.population >= 25000000;