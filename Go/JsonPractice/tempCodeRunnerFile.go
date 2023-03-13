countStud, countOcenk := len(jsin.Students), 0;
	for i := 0; i < len(jsin.Students); i++ {
		countOcenk += len(jsin.Students[i].Rating); 
	}

	sr := float64(countOcenk) / float64(countStud);