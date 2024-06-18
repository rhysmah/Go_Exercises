import React, {useEffect, useState} from 'react';

// Call backend API
const ContentBox = () => {
    const [data, setData] = useState(null);
    const [error, setError] = useState(null);

    useEffect(() => {
        console.log("Component mounted; fetching data...")
        fetch('http://localhost:8080/story')
        .then(response => {
            console.log("Received response: ", response)

            if (!response.ok) {
                throw new Error('Failed to fetch data')
            }
            console.log(response.json())
            return response.json()
        })
        .then(data => setData(data))
        .catch(error => setError(error));
    }, []);

    if (error) {
        return <p>{error.message}</p>
    }

    if (!data) {
        return <p>Loading...</p>
    }

    return (
        <div className="content-box">
            <h1>Content Box</h1>
        </div>
    )
}


export default ContentBox;