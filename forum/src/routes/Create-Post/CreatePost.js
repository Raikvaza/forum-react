import React from 'react';
import Header from '../../components/Header/Header';
import Body from '../../components/Body/Body';

// import { useLocation } from "react-router-dom";


const CreatePost = () => {
// const {username} = props;
// const location = useLocation();
// const { state } = location;
// const username = state.username;


    return (
        <div>
        <Header />
        <Body createPost={true}/>
        <div>
        {/* <h2> Hello {username}</h2> */}

        </div>
        </div>
    );
};

export default CreatePost;