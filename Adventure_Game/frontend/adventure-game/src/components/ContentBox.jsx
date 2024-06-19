import React, {useEffect, useState} from 'react';

const ContentBox = () => {
    const [data, setData] = useState(null);
    const [error, setError] = useState(null);

    useEffect( () => {
        console.log("Component mounted. Fetching data...")
        fetch('http://localhost:8080/story')
        .then(response => {
            if (!response.ok) {
                setError('Failed to fetch data')
            }
            return response.json()
        })
        .then(data => setData(data))
        .catch(error => setError(error))
    })

    if (error) return <p>Error: {error}</p>
    if (!data) return <p>Loading...</p>

    return (
        <>
            <div className="content-box">
                <h1>{data.intro.title}</h1>
                <p>{data.intro.story}</p>
            </div>

            <div className="links-box">
                <ul>
                    {data.intro.options.map((option, index) => (
                        <li key={index}><a href={option.link}>{option.text}</a></li>
                    ))}
                </ul>
            </div>

        </>
    )
}


export default ContentBox;