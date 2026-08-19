import { Link } from "react-router-dom"

export default function Landing() {
    return (
        <>
        <div className="min-h-screen bg-yellow-500">
            <h1 className="text-3xl font-bold text-center">Welcome to the Landing Page</h1>
            <div className="text-right mr-10 mt-5">
                <div><Link to="/login" >
                <button type="button" className=" hover:underline ">login </button>
                </Link>
                </div>
                <div>
                <Link to="/Register">
                <button className=" hover:underline"> register</button>
                </Link>
                </div>
            </div>
        </div>


        </>
    )
}

