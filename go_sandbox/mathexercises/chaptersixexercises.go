package mathexercises

//6.1 The heavy Pill: You have 20 bottles of pills. 19 bottles have 1.0 gram pills, but 1 has pills of weight 1.1 grams. Given a scale that provides exact measurement, how would you find the heavy bottle? You can only use the scale once.

//take 1 pill from bottle 1, 2 pills from bottle 2, etc.
//the overweight will be .1*the heavy bottle.

//6.2 Basketball: play one of two games:
//1: One shot to make the hoop
//2: three shot to make two shots.
//P==probability
//What values of P are optimal for each game?
//P(a and b)=P(b given a)P(a)
//P(a or B)=P(a) + P(b) -P (a and b)
//Game 1:P=P
//Game 2: P(a and b) or P(b and c) or P(a and c)
//P(Game 2)=(P^2 OR P^2) OR P^2
//P(Game 2)=(2P^2-P^4) OR P^2
//P(Game 2)=()(3P^2)-(P^4))-((2P^4)-(P^6))
//play game 1 when P<.389391, play P when P>.389391

//wrong lol, didnt account for missing the third shots
