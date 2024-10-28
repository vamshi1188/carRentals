import React, { useState } from 'react';
import axios from '../services/api';

function PaymentForm({ amount }) {
  const [clientSecret, setClientSecret] = useState('');

  const handlePayment = async () => {
    try {
      const response = await axios.post('/process-payment', { amount });
      setClientSecret(response.data.client_secret);
      // Initiate payment with clientSecret
    } catch (error) {
      console.error('Payment failed:', error);
    }
  };

  return (
    <div>
      <button onClick={handlePayment}>Pay ${amount / 100}</button>
      {clientSecret && <p>Payment ready!</p>}
    </div>
  );
}

export default PaymentForm;
