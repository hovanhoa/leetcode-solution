DELETE P1 FROM Person AS P1  
INNER JOIN Person AS P2   
WHERE P2.id < P1.id AND P1.email = P2.email;  