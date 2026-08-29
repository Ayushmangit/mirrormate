import { Link, useNavigate } from "react-router-dom";
import { useForm } from "react-hook-form";
import { useAppDispatch, useAppSelector } from "../../app/hooks";
import { registerUser } from "../../features/auth/authThunk.ts";

type RegisterForm = {
	username: string;
	email: string;
	password: string;
	confirmPassword: string;
};

function RegisterPage() {
	const dispatch = useAppDispatch();
	const navigate = useNavigate();

	const { loading, error } = useAppSelector((state) => state.auth);

	const {
		register,
		handleSubmit,
		watch,
		formState: { errors },
	} = useForm<RegisterForm>();

	const password = watch("password");

	const onSubmit = async ({ username, email, password }: RegisterForm) => {
		const payload = {
			username: username,
			email: email,
			password: password,
		};

		const result = await dispatch(registerUser(payload));

		if (registerUser.fulfilled.match(result)) {
			navigate("/login");
		}
	};

	return (
		<div className="flex min-h-screen items-center justify-center bg-gray-50 px-6">
			<div className="w-full max-w-md">
				<div className="mb-8 text-center">
					<Link
						to="/"
						className="text-2xl font-bold tracking-tight text-gray-900"
					>
						BoardGo
					</Link>

					<h1 className="mt-6 text-3xl font-bold tracking-tight text-gray-900">
						Create your account
					</h1>

					<p className="mt-2 text-sm text-gray-600">
						Start creating your visual boards.
					</p>
				</div>

				<form
					onSubmit={handleSubmit(onSubmit)}
					className="space-y-5 rounded-xl border border-gray-200 bg-white p-8 shadow-sm"
				>
					<div>
						<label
							htmlFor="username"
							className="mb-2 block text-sm font-medium text-gray-700"
						>
							Username
						</label>

						<input
							id="username"
							type="text"
							autoComplete="username"
							placeholder="ayushman"
							{...register("username", {
								required: "Username is required",

								minLength: {
									value: 3,
									message: "Username must be at least 3 characters",
								},

								maxLength: {
									value: 50,
									message: "Username must be at most 50 characters",
								},
							})}
							className="w-full rounded-lg border border-gray-300 px-3 py-2 outline-none transition focus:border-gray-900"
						/>

						{errors.username && (
							<p className="mt-2 text-sm text-red-600">
								{errors.username.message}
							</p>
						)}
					</div>

					<div>
						<label
							htmlFor="email"
							className="mb-2 block text-sm font-medium text-gray-700"
						>
							Email
						</label>

						<input
							id="email"
							type="email"
							autoComplete="email"
							placeholder="you@example.com"
							{...register("email", {
								required: "Email is required",

								pattern: {
									value: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
									message: "Enter a valid email address",
								},
							})}
							className="w-full rounded-lg border border-gray-300 px-3 py-2 outline-none transition focus:border-gray-900"
						/>

						{errors.email && (
							<p className="mt-2 text-sm text-red-600">
								{errors.email.message}
							</p>
						)}
					</div>

					<div>
						<label
							htmlFor="password"
							className="mb-2 block text-sm font-medium text-gray-700"
						>
							Password
						</label>

						<input
							id="password"
							type="password"
							autoComplete="new-password"
							placeholder="Create a password"
							{...register("password", {
								required: "Password is required",

								minLength: {
									value: 8,
									message: "Password must be at least 8 characters",
								},
							})}
							className="w-full rounded-lg border border-gray-300 px-3 py-2 outline-none transition focus:border-gray-900"
						/>

						{errors.password && (
							<p className="mt-2 text-sm text-red-600">
								{errors.password.message}
							</p>
						)}
					</div>

					<div>
						<label
							htmlFor="confirmPassword"
							className="mb-2 block text-sm font-medium text-gray-700"
						>
							Confirm password
						</label>

						<input
							id="confirmPassword"
							type="password"
							autoComplete="new-password"
							placeholder="Confirm your password"
							{...register("confirmPassword", {
								required: "Please confirm your password",

								validate: (value) =>
									value === password || "Passwords do not match",
							})}
							className="w-full rounded-lg border border-gray-300 px-3 py-2 outline-none transition focus:border-gray-900"
						/>

						{errors.confirmPassword && (
							<p className="mt-2 text-sm text-red-600">
								{errors.confirmPassword.message}
							</p>
						)}
					</div>

					{error && (
						<div className="rounded-lg bg-red-50 px-3 py-2 text-sm text-red-600">
							{error}
						</div>
					)}

					<button
						type="submit"
						disabled={loading}
						className="w-full rounded-lg bg-gray-900 px-4 py-3 font-medium text-white transition hover:bg-gray-700 disabled:cursor-not-allowed disabled:opacity-60"
					>
						{loading ? "Creating account..." : "Create account"}
					</button>
				</form>

				<p className="mt-6 text-center text-sm text-gray-600">
					Already have an account?{" "}
					<Link to="/login" className="font-medium text-gray-900 underline">
						Sign in
					</Link>
				</p>
			</div>
		</div>
	);
}

export default RegisterPage;
