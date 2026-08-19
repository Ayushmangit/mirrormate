import React from "react";
import { useState, type ChangeEvent } from 'react';



const Register = () => {
  const [formData, setFormData] = useState({
    fullName: '',
    emailAddress: '',
    password: '',
    confirmPassword: '',
    role: 'user',
  });
    const handleChange =
     (e: ChangeEvent<HTMLInputElement |HTMLSelectElement>) => {
          setFormData({
              ...formData,
              [e.target.name]: e.target.value,
          });
      };

  return (
    <div>
      <h1>Create Account</h1>
<form>
{/*Full Name*/}
<label>Full Name</label>
<input type="text" id="fullName"
name="fullName"
placeholder="Enter your full name"
 onChange={handleChange}
 />

{/*Email Address*/}
<label>Email Address</label>
<input type="text" id="emailAddress" name="emailAddress"
placeholder="Enter your Email Address"
 onChange={handleChange}/>

{/*Password*/}
<label>Password</label>
<input type="password" id="password" name="password" placeholder="Enter your Password"
 onChange={handleChange}
 />

{/*Confirm Password*/}
<label>Confirm Password</label>
<input type="password" id="confirmPassword" name="confirmPassword"
placeholder="Enter your Password"
 onChange={handleChange}/>

{/*Role*/}
<label>Role</label>
<select id="role" name="role" onChange={handleChange}>

  <option value="user">User</option>

</select>

{/*Submit*/}
<button type="submit">Register</button>
      </form>
    </div>
  );
};

export default Register
