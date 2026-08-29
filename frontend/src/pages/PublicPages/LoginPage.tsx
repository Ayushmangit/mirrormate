import { Link, useNavigate } from "react-router-dom";
import { useForm } from "react-hook-form";
import { useAppDispatch, useAppSelector } from "../../app/hooks";
import { loginUser } from "../../features/auth/authThunk";

type LoginForm = {
	email: string;
	password: string;
};

function LoginPage() {
	const dispatch = useAppDispatch();
	const navigate = useNavigate();

	const { loading, error } = useAppSelector((state) => state.auth);

	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm<LoginForm>();

	const onSubmit = async (data: LoginForm) => {
		const result = await dispatch(loginUser(data));

		if (loginUser.fulfilled.match(result)) {
			navigate("/boards");
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
						Welcome back
					</h1>

					<p className="mt-2 text-sm text-gray-600">
						Sign in to continue to your boards.
					</p>
				</div>

				<form
					onSubmit={handleSubmit(onSubmit)}
					className="space-y-5 rounded-xl border border-gray-200 bg-white p-8 shadow-sm"
				>
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
							autoComplete="current-password"
							placeholder="Enter your password"
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
						{loading ? "Signing in..." : "Sign in"}
					</button>
				</form>

				<p className="mt-6 text-center text-sm text-gray-600">
					Don&apos;t have an account?{" "}
					<Link to="/register" className="font-medium text-gray-900 underline">
						Create one
					</Link>
				</p>
			</div>
		</div>
	);
}

export default LoginPage;
