import React, { useState } from 'react';
import { registerUser } from '../api';

const Register = ({ goToLogin }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [msg, setMsg] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setMsg('');
    try {
      await registerUser({ username, password });
      setMsg('Registration successful. You can now log in.');
      setUsername('');
      setPassword('');
      if (goToLogin) goToLogin();
    } catch {
      setError('Registration failed. Try a different username.');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Register</h2>
      {error && <p style={{color:'red'}}>{error}</p>}
      {msg && <p style={{color:'green'}}>{msg}</p>}
      <input placeholder="Username" value={username} onChange={e => setUsername(e.target.value)} required />
      <input placeholder="Password" value={password} type="password" onChange={e => setPassword(e.target.value)} required />
      <button type="submit">Register</button>
    </form>
  );
};

export default Register;
