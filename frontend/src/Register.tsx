import { useForm } from "react-hook-form"

type RegisterFormData = {
	username: string
	email: string
	password: string
	confirmPassword: string
}

export default function RegisterUser() {
	const {
		register,
		handleSubmit,
		watch,
		formState: { errors },
	} = useForm<RegisterFormData>()

	const password = watch("password")

	function onSubmit(data: RegisterFormData) {
		console.log("Registration data:", data)
		alert("Registration form is valid!")
	}

	return (
		<div className="min-h-screen bg-gray-900 flex items-center justify-center p-4">
			<div className="w-full max-w-md rounded-3xl bg-white p-8 shadow-lg sm:p-10">

					{/* Heading */}
					<div className="mb-8 text-center">
						<h1 className="text-3xl font-extrabold sm:text-4xl"
							style={{ color:"#000000", opacity:1}}
							>
							Create your account
						</h1>

						<p className="mt-3 text-gray-600">
							Join MirrorMate today.
						</p>
					</div>

					<form
						onSubmit={handleSubmit(onSubmit)}
						className="space-y-5"
					>

						{/* Username */}
						<div>
							<label
								htmlFor="username"
								className="mb-2 block text-sm font-semibold text-gray-700"
							>
								Username
							</label>

							<input
								id="username"
								type="text"
								placeholder="Enter your username"
								{...register("username", {
									required: "Username is required",
									minLength: {
										value: 3,
										message:
											"Username must be at least 3 characters",
									},
								})}
								className="w-full rounded-xl border border-gray-300 bg-gray-50 px-4 py-3 text-gray-900 outline-none focus:border-blue-500 focus:ring-2 focus:ring-blue-100"
							/>

							{errors.username && (
								<p className="mt-1 text-sm text-red-500">
									{errors.username.message}
								</p>
							)}
						</div>

						{/* Email */}
						<div>
							<label
								htmlFor="email"
								className="mb-2 block text-sm font-semibold text-gray-700"
							>
								Email Address
							</label>

							<input
								id="email"
								type="email"
								placeholder="Enter your email address"
								{...register("email", {
									required: "Email is required",
									pattern: {
										value:
											/^[^\s@]+@[^\s@]+\.[^\s@]+$/,
										message:
											"Enter a valid email address",
									},
								})}
								className="w-full rounded-xl border border-gray-300 bg-gray-50 px-4 py-3 text-gray-900 outline-none focus:border-blue-500 focus:ring-2 focus:ring-blue-100"
							/>

							{errors.email && (
								<p className="mt-1 text-sm text-red-500">
									{errors.email.message}
								</p>
							)}
						</div>

						{/* Password */}
						<div>
							<label
								htmlFor="password"
								className="mb-2 block text-sm font-semibold text-gray-700"
							>
								Password
							</label>

							<input
								id="password"
								type="password"
								placeholder="Enter your password"
								{...register("password", {
									required: "Password is required",
									minLength: {
										value: 4,
										message:
											"Password must be at least 4 characters",
									},
									pattern: {
										value:
											/^(?=.*[A-Za-z])(?=.*\d).{4,}$/,
										message:
											"Password must contain at least one letter and one number",
									},
								})}
								className="w-full rounded-xl border border-gray-300 bg-gray-50 px-4 py-3 text-gray-900 outline-none focus:border-blue-500 focus:ring-2 focus:ring-blue-100"
							/>

							<p className="mt-1 text-xs text-gray-500">
								Minimum 4 characters with at least one letter
								and one number.
							</p>

							{errors.password && (
								<p className="mt-1 text-sm text-red-500">
									{errors.password.message}
								</p>
							)}
						</div>

						{/* Confirm Password */}
						<div>
							<label
								htmlFor="confirmPassword"
								className="mb-2 block text-sm font-semibold text-gray-700"
							>
								Confirm Password
							</label>

							<input
								id="confirmPassword"
								type="password"
								placeholder="Confirm your password"
								{...register("confirmPassword", {
									required:
										"Please confirm your password",
									validate: (value) =>
										value === password ||
										"Passwords do not match",
								})}
								className="w-full rounded-xl border border-gray-300 bg-gray-50 px-4 py-3 text-gray-900 outline-none focus:border-blue-500 focus:ring-2 focus:ring-blue-100"
							/>

							{errors.confirmPassword && (
								<p className="mt-1 text-sm text-red-500">
									{errors.confirmPassword.message}
								</p>
							)}
						</div>

						{/* Submit */}
						<button
							type="submit"
							className="w-full rounded-xl bg-blue-600 px-6 py-3 font-semibold text-white transition hover:bg-blue-700"
						>
							Create Account
						</button>
					</form>

					<div className="mt-8 text-center text-sm text-gray-600">
						Already have an account?
					</div>
				</div>
			</div>
	)
}