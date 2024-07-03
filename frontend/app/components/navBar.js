import React, { useState, useEffect } from 'react';
import {Modal, Button, Form, Input, message} from 'antd';
import UserService from '../../src/service/UserService';

const userService = new UserService();

const Navbar = () => {
    const [userInfo, setUserInfo] = useState(null);
    const [modalVisible, setModalVisible] = useState(false);
    const [loginMode, setLoginMode] = useState(true); // 控制显示登录还是注册表单

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            userService.GetUserInfoByToken(token)
                .then(data => {
                    if (data.code === 0) {
                        setUserInfo(data.data);
                    } else {
                        console.log(data.message);
                    }
                })
                .catch(error => {
                    console.error('获取用户信息出错:', error);
                });
        }
    }, []);

    const handleModalOpen = () => {
        setModalVisible(true);
    };

    const handleModalClose = () => {
        setModalVisible(false);
    };

    const handleSignOut = () => {
        userService.signOut();
        setUserInfo(null);
        handleModalClose();
    };

    const handleSwitchMode = () => {
        setLoginMode(prevMode => !prevMode);
    };

    const onFinish = async (values) => {
        const { username, password, email } = values;
        if (loginMode) {
            // 登录
            const result = await userService.signIn(username, password);
            if (result.code === 0) {
                const token = result.data.token;
                localStorage.setItem('token', token);
                const userInfoResult = await userService.GetUserInfoByToken(token);
                if (userInfoResult.code === 0) {
                    setUserInfo(userInfoResult.data);
                } else {
                    console.log(userInfoResult.message);
                }
                message.success("登录成功！");
                handleModalClose();
            } else {
                message.error(result.message);
            }
        } else {
            // 注册
            const result = await userService.signUp(username, email, password);
            if (result.code === 0) {
                message.success("注册成功！");
                handleSwitchMode();
            } else {
                message.error(result.message);
            }
        }
    };

    return (
        <div className="fixed top-0 left-0 w-full h-20 bg-blue-500 bg-opacity-60 backdrop-filter backdrop-blur-md shadow-md z-50 flex items-center justify-between px-4">
            <div className="flex items-center border-r-2 border-gray-300 pr-4">
                <img
                    src="/usst.png" // 替换为你的图片路径
                    alt="Logo"
                    className="mr-4"
                />
                <span className="text-white text-xl font-semibold">
                    上理的课
                </span>
            </div>
            <div className="flex items-center pr-32">
                <img
                    src="/avatar.jpg" // 替换为你的头像路径
                    alt="User Avatar"
                    className="h-12 w-12 rounded-full border-2 border-white mr-4 cursor-pointer"
                    onClick={handleModalOpen}
                />
                <span className="text-white text-xl font-semibold cursor-pointer" onClick={handleModalOpen}>
                    {userInfo ? userInfo.username : '未登录'}
                </span>
            </div>

            <Modal
                title={loginMode ? '用户登录' : '用户注册'}
                visible={modalVisible}
                onCancel={handleModalClose}
                footer={null} // 禁用默认的底部按钮
            >
                {userInfo ? (
                    <div>
                        <p>用户名: {userInfo.username}</p>
                        <p>邮箱: {userInfo.email}</p>
                        <Button onClick={handleSignOut}>退出登录</Button>
                    </div>
                ) : (
                    <Form
                        name="loginForm"
                        initialValues={{ remember: true }}
                        onFinish={onFinish}
                    >
                        <Form.Item
                            name="username"
                            rules={[{ required: true, message: '请输入用户名' }]}
                        >
                            <Input placeholder="用户名" />
                        </Form.Item>
                        {!loginMode && (
                            <Form.Item
                                name="email"
                                rules={[{ required: true, message: '请输入邮箱' }]}
                            >
                                <Input placeholder="邮箱" />
                            </Form.Item>
                        )}
                        <Form.Item
                            name="password"
                            rules={[{ required: true, message: '请输入密码' }]}
                        >
                            <Input.Password placeholder="密码" />
                        </Form.Item>
                        <Form.Item>
                            <Button type="primary" htmlType="submit">
                                {loginMode ? '登录' : '注册'}
                            </Button>
                            <Button onClick={handleSwitchMode}>
                                切换到 {loginMode ? '注册' : '登录'}
                            </Button>
                        </Form.Item>
                    </Form>
                )}
            </Modal>
        </div>
    );
};

export default Navbar;
