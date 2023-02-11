import React, {useEffect, useState} from 'react';
import Header from '../../components/Header/Header';
import Body from '../../components/Body/Body';
import './HomePage.css'
// import { json } from 'react-router-dom';

const HomePage = () => {

const [posts, setPosts] = useState([{}]);
const [isAuth, setIsAuth] = useState(false);    

useEffect (() => {
    const fetchData = async () => {
      
      await fetch(`http://localhost:8080/api/checkUser`, {
        headers: {
          'Accept': 'application/json ',
          'Credentials': 'include'
        },
        method: "GET",
        credentials: 'include',
      }).then((r) => {
          if(r.ok){
            setIsAuth (true);
            return r.json();
          } else if (r.status === 401){
            setIsAuth(false)
            return null
          } else {
            throw new Error("Server error")
          }
        }
        )
      
      await fetch(`http://localhost:8080/api/home`, {
          headers: {
            'Accept': 'application/json ',
            'Credentials': 'include'
          },
          method: "GET",
          credentials: 'include',
        }).then((r) => r.json())
        .then((data) => {
            if (data !== null){
              setPosts(data)
            } else {
              console.log("No posts in Homepage");
            }
        });
      };
      fetchData();
    },[]); 

    return (
        <div>
        <Header status = {isAuth}/>
        <Body createPost = {false} posts={posts}/>
        <div>
        </div>
        </div>
    );
};

export default HomePage;