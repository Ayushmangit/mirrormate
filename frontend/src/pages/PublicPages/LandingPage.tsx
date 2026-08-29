import { Link } from "react-router-dom";

function LandingPage() {
	return (
		<div className="min-h-screen bg-white text-gray-900">
			<Navbar />
			<main>
				<Hero />
			</main>
		</div>
	);
}

function Navbar() {
	return (
		<header className="border-b border-gray-200">
			<nav className="mx-auto flex max-w-7xl items-center justify-between px-6 py-4">
				<Link to="/" className="text-xl font-bold tracking-tight">
					BoardGo
				</Link>

				<div className="flex items-center gap-3">
					<Link
						to="/login"
						className="rounded-lg px-4 py-2 text-sm font-medium text-gray-700 transition hover:bg-gray-100"
					>
						Login
					</Link>

					<Link
						to="/register"
						className="rounded-lg bg-gray-900 px-4 py-2 text-sm font-medium text-white transition hover:bg-gray-700"
					>
						Get Started
					</Link>
				</div>
			</nav>
		</header>
	);
}

function Hero() {
	return (
		<section className="mx-auto flex min-h-[calc(100vh-73px)] max-w-7xl items-center px-6">
			<div className="max-w-3xl">
				<p className="mb-4 text-sm font-semibold uppercase tracking-wider text-gray-500">
					Visual collaboration
				</p>

				<h1 className="text-5xl font-bold tracking-tight sm:text-6xl">
					Turn your ideas into
					<span className="block text-gray-500">visual boards.</span>
				</h1>

				<p className="mt-6 max-w-2xl text-lg leading-8 text-gray-600">
					Create diagrams, sketch ideas, organize thoughts, and build visual
					workflows on an infinite canvas.
				</p>

				<div className="mt-8 flex items-center gap-4">
					<Link
						to="/register"
						className="rounded-lg bg-gray-900 px-6 py-3 font-medium text-white transition hover:bg-gray-700"
					>
						Start Drawing
					</Link>

					<Link
						to="/login"
						className="rounded-lg border border-gray-300 px-6 py-3 font-medium text-gray-700 transition hover:bg-gray-50"
					>
						Sign In
					</Link>
				</div>
			</div>
		</section>
	);
}

export default LandingPage;
