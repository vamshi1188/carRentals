import React, { useEffect, useState } from 'react';
import axios from '../services/api';

function CarList() {
  const [cars, setCars] = useState([]);

  useEffect(() => {
    axios.get('/cars')
      .then(response => setCars(response.data))
      .catch(error => console.error('Error fetching cars:', error));
  }, []);

  return (
    <div>
      {cars.map(car => (
        <div key={car.id}>
          <img src={car.image} alt={`${car.make} ${car.model}`} />
          <h2>{car.make} {car.model}</h2>
          <p>Year: {car.year}</p>
          <p>Price: ${car.price}</p>
        </div>
      ))}
    </div>
  );
}

export default CarList;
