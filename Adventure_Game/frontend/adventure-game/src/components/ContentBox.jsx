import React, {useEffect, useState} from 'react';

const ContentBox = () => {
    const [error, setError]           = useState(null);
    const [storyData, setStoryData]   = useState(null);
    const [currentArc, setCurrentArc] = useState(null);
    

    useEffect( () => {
        fetch('http://localhost:8080/story')
        .then(response => {
            if (!response.ok) throw new Error("Failed to fetch data")
            return response.json()
        })
        .then(response => {
            setStoryData(response)        // Stores all story data
            console.log(response)
            setCurrentArc(response.intro) // Stores intro arc
        })
        .catch(error => setError(error))
    }, []);


    if (error) return <p>Error: {error}</p>
    if (!storyData) return <p>Loading...</p>


    const {title, story, options} = currentArc;


    return (
        <>
             <div className="content-box">
                <h1>{title}</h1>
                {story.map((paragraph, index) => (
                    <p key={index}>{paragraph}</p>
                ))}
            </div>

            <div className="links-box">
                <ul>
                    {options.map((option, index) => (
                        <li key={index}>
                            <a href={option.link} 
                            onClick={(e) => {
                                e.preventDefault()
                                setCurrentArc(storyData[option.arc])
                            }
                            }>
                                {option.text}
                            </a>
                        </li>
                    ))}
                </ul>
            </div>
        </>
    )
}

export default ContentBox;