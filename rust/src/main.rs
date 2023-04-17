use std::{
    fs,
    io::{Read, Write},
    os::unix::net::{UnixListener, UnixStream},
    time::Duration,
};

const PATH: &str = "/tmp/ipc-test.sock";

fn main() {
    // let file = fs::OpenOptions::new()
    //     .create_new(true)
    //     .write(true)
    //     .open(PATH)
    //     .unwrap();

    // drop(file);

    // let listener = UnixListener::bind(PATH).unwrap();

    // let (mut stream, _addr) = listener.accept().unwrap();

    loop {
        let mut stream = UnixStream::connect(PATH).unwrap();

        let msg = "hello world";

        println!("[MESSAGE ] {msg}");

        stream.write_all(msg.as_bytes()).unwrap();
        let mut response = String::new();

        stream.read_to_string(&mut response).unwrap();

        println!("[RESPONSE] {response}");

        std::thread::sleep(Duration::from_secs(6));
    }
}
