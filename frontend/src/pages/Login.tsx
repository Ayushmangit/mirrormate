import { useState } from "react";
import type { ChangeEvent, FormEvent } from "react";

const EMAIL_RE = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

export default function Login() {
    const [formData, setFormData] = useState({
        email: "",
        password: "",
    });

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value,
        });
    };

    const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        console.log(formData);

        if (!EMAIL_RE.test(formData.email)) {
            console.log("Invalid email");
            return;
        }

        console.log("Email is valid");

        
    };  

    return (
        <>
            <div className="min-h-screen">
                <h1>Welcome to the Login Page</h1>

                <form
                    onSubmit={handleSubmit}
                    className="flex flex-col gap-4"
                >
                    <div>
                        <label htmlFor="email">
                            Email:
                        </label>
                        <br />

                        <input
                            type="email"
                            name="email"
                            id="email"
                            value={formData.email}
                            onChange={handleChange}
                            className="border-2 border-black"
                        />
                    </div>

                    <div>
                        <label htmlFor="password">
                            Password:
                        </label>
                        <br />

                        <input
                            type="password"
                            name="password"
                            id="password"
                            value={formData.password}
                            onChange={handleChange}
                            className="border-2 border-black"
                        />
                    </div>

                    <button type="submit">
                        Login
                    </button>
                </form>
            </div>
        </>
    );
}