import React, {useState, useEffect} from 'react';
import Header from '../../components/Header/Header';
import Body from '../../components/Body/Body';
import { useNavigate } from 'react-router-dom';


const CreatePost = () => {

const [user, setUser] = useState();    
const navigate = useNavigate();
const [isLoading, setIsLoading] = useState(true);
useEffect (() => {
    (async () => {
      await fetch(`http://localhost:8080/api/checkUser`, {
        headers: {
          'Accept': 'application/json ',
          'Credentials': 'include'
        },
        method: "GET",
        credentials: 'include',
      }).then((r) => {
            if(r.ok){
                return r.json();
            } else if (r.status === 401){
                console.log("Can't go to createPost if unauth");
                navigate("/")
            } else {
                throw new Error("Server error")
            }
        }).then((data) => {
          if (data!== null){
            setUser(data.Username);
          } else {
            console.log("Error in Create Post");
          }
        });
      })();
    },[]);
    useEffect(() => {
        if (user!== undefined){
            console.log(user);
            setIsLoading(false);
        }
    }, [user]) 

    return isLoading ? (
        <div>Loading...</div>
      ) : (
        <div>
        <Header />
        <Body createPost={true} username={user}/>
        </div>
      );
};

export default CreatePost;