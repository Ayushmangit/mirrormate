    import { Link, useNavigate } from "react-router-dom";
    import { useForm } from "react-hook-form";
    import {useAppDispatch, useAppSelector} from "../hooks";
    import { loginUser } from "../features/auth/Auththunk";    
import type { LoginUser } from "../types/user";

    export default function Login() {
        const dispatch = useAppDispatch();
        const navigate = useNavigate();
        const { register, handleSubmit } = useForm<LoginUser>();//userForm(loginUser)();
        const { isLoading, error } = useAppSelector(
            (state) => state.auth);
        
        async function onSubmit(data: LoginUser) {
            try {
                await dispatch(loginUser(data)).unwrap();
                navigate("/home");
            } catch (error) {
                console.error("Login failed:", error);
            }
        };

        return (
            <div >
                <div >
                    <h2 >Login</h2>
                    
                    <form onSubmit={handleSubmit(onSubmit)}>
                        <div>
                            <label >
                                Email
                            </label>
                            <input
                                id="email"
                                type="email"
                                {...register("email", { required: true })}
                                />
                        </div>
                        <div>
                            <label >
                                Password
                            </label>
                            <input
                                id="password"
                                type="password"
                                {...register("password", { required: true })}
                                />
                        </div>
                        <button
                            type="submit"
                            disabled={isLoading}
                            
                        >
                            {isLoading ? "Logging in..." : "Login"}
                        </button>
                    </form>

                    <p className="text-sm text-center text-gray-600">
                        Don't have an account?{" "}
                        <Link to="/signup" >
                            Sign up
                        </Link>
                    </p>
                </div>
            </div>
        );
        
    }