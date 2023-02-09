import React, {useEffect, useState} from 'react';
import Header from '../../components/Header/Header';
import Body from '../../components/Body/Body';
import './HomePage.css'
import { json } from 'react-router-dom';

const HomePage = () => {
const [posts, setPosts] = useState([{}]);
useEffect (() => {
    // (async() => {
    //     await fetch(`http://localhost:8080/api/checkUserByToken`, 
    //     {
    //         headers: {
    //             'Accept': 'application/json ',
    //             'Credentials': 'include'
    //         },
    //         method: "GET",
    //         credentials: 'include',
    //     }).then((r) => {
    //         console.log(r);
    //     })
    // })()
    
    // (async() => {
    //     await fetch(`http://localhost:8080/api/getAllpost`, 
    //     {
    //         headers: {
    //             'Accept': 'application/json ',
    //             'Content-type': 'text/plain',
    //             'Credentials': 'include'
    //         },
    //         method: "GET",
    //         credentials: 'include',
    //     }).then((r) => {
    //         console.log(r);
    //     })
    // })()

    const fetchData = async () => {
        await fetch(`http://localhost:8080/api/checkUserByToken`, {
          headers: {
            'Accept': 'application/json ',
            'Credentials': 'include'
          },
          method: "GET",
          credentials: 'include',
        }).then((r) => r.json())
        .then((data) => {
            console.log(data);
            //Store this data in the local storage
        });
    
        await fetch(`http://localhost:8080/api/getAllpost`, {
          headers: {
            'Accept': 'application/json ',
            'Content-type': 'text/plain',
            'Credentials': 'include'
          },
          method: "GET",
          credentials: 'include',
        }) .then(response => response.json())
        .then(data => {setPosts(data)})
        .catch(error => console.error(error));
      };
      fetchData();
      
},[]);

    return (
        <div>
        <Header />
        <Body posts={posts}/>
        <div>
        {/* <h2> Hello {username}</h2> */}
        </div>
        </div>
    );
};

export default HomePage;