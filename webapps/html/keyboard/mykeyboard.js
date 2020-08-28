let Keyboard = window.SimpleKeyboard.default;

let commonKeyboardOptions = {
    onChange: input => onChange(input),
    onKeyPress: button => onKeyPress(button),
    theme: "simple-keyboard hg-theme-default hg-layout-default",
    physicalKeyboardHighlight: true,
    syncInstanceInputs: true,
    mergeDisplay: true,
    debug: true,
    //useTouchEvents: true
    //newLineOnEnter: true,
    //disableButtonHold: true,
    //onKeyReleased: button => onKeyReleased(button),

};

let keyboard = new Keyboard(".simple-keyboard-main", {
    ...commonKeyboardOptions,
    /**
     * Layout by:
     * Sterling Butters (https://github.com/SterlingButters)
     */
    layout: {
        default: [
            "{escape} {f1} {f2} {f3} {f4} {f5} {f6} {f7} {f8} {f9} {f10} {f11} {f12}",
            "` 1 2 3 4 5 6 7 8 9 0 - = {backspace}",
            "{tab} q w e r t y u i o p [ ] \\",
            "{capslock} a s d f g h j k l ; ' {enter}",
            "{shiftleft} z x c v b n m , . / {shiftright}",
            "{controlleft} {altleft} {metaleft} {space} {metaright} {altright}"
        ],
        shift: [
            "{escape} {f1} {f2} {f3} {f4} {f5} {f6} {f7} {f8} {f9} {f10} {f11} {f12}",
            "~ ! @ # $ % ^ & * ( ) _ + {backspace}",
            "{tab} Q W E R T Y U I O P { } |",
            '{capslock} A S D F G H J K L : " {enter}',
            "{shiftleft} Z X C V B N M < > ? {shiftright}",
            "{controlleft} {altleft} {metaleft} {space} {metaright} {altright}"
        ]
    },
    display: {
        "{escape}": "esc ⎋",
        "{tab}": "tab ⇥",
        "{backspace}": "backspace ⌫",
        "{enter}": "enter ↵",
        "{capslock}": "caps lock ⇪",
        "{shiftleft}": "shift ⇧",
        "{shiftright}": "shift ⇧",
        "{controlleft}": "ctrl ⌃",
        "{controlright}": "ctrl ⌃",
        "{altleft}": "alt ⌥",
        "{altright}": "alt ⌥",
        "{metaleft}": "cmd ⌘",
        "{metaright}": "cmd ⌘"
    }
});

let ctrlLock = false;
function onChange(input) {
    //document.querySelector(".input").value = input;
    keyboard.setInput(input);

    console.log("Input changed", input);
}

function onKeyPress(button) {
    console.log("Button pressed", button);

    /**
     * If you want to handle the shift and caps lock buttons
     */
    if (
        button === "{shift}" ||
        button === "{shiftleft}" ||
        button === "{shiftright}" ||
        button === "{capslock}"
    ){
        handleShift();
    }else if(
        button === "{controlleft}"
    ){
        handleCtrl();
    }else{
        handleOther(button)
    }
}

function handleOther(button) {
    let code = buttonToCharCode(button);
    if(socket){
        socket.send(JSON.stringify({ type: "stdin", data: utf8_to_b64(code) }));
    }

    releasedCtrl();
}

function handleCtrl() {
    let b = $('[data-skbtn="{controlleft}"]');
    b.addClass("ctrl-active");
    ctrlLock = true;
}
function releasedCtrl() {
    let b = $('[data-skbtn="{controlleft}"]');
    b.removeClass("ctrl-active");
    ctrlLock = false;
}

function handleShift() {
    let currentLayout = keyboard.options.layoutName;
    let shiftToggle = currentLayout === "default" ? "shift" : "default";

    keyboard.setOptions({
        layoutName: shiftToggle
    });
}

function buttonToCharCode(button) {

    if(ctrlLock){
        switch (button) {
            case 'a': return String.fromCharCode(1);
            case 'b': return String.fromCharCode(2);
            case 'c': return String.fromCharCode(3);
            case 'd': return String.fromCharCode(4);
            case 'e': return String.fromCharCode(5);
            case 'f': return String.fromCharCode(6);
            case 'g': return String.fromCharCode(7);
            case 'h': return String.fromCharCode(8);
            case 'i': return String.fromCharCode(9);
            case 'j': return String.fromCharCode(10);
            case 'k': return String.fromCharCode(11);
            case 'l': return String.fromCharCode(12);
            case 'm': return String.fromCharCode(13);
            case 'n': return String.fromCharCode(14);
            case 'o': return String.fromCharCode(15);
            case 'p': return String.fromCharCode(16);
            case 'q': return String.fromCharCode(17);
            case 'r': return String.fromCharCode(18);
            case 's': return String.fromCharCode(19);
            case 't': return String.fromCharCode(20);
            case 'u': return String.fromCharCode(21);
            case 'v': return String.fromCharCode(22);
            case 'w': return String.fromCharCode(23);
            case 'x': return String.fromCharCode(24);
            case 'y': return String.fromCharCode(25);
            case 'z': return String.fromCharCode(26);
        }
    }
    switch (button) {
        case '{escape}': return String.fromCharCode(27);
        case '{backspace}': return String.fromCharCode(8);
        case '{tab}': return String.fromCharCode(9);
        case '{enter}': return String.fromCharCode(13);
        case '{space}': return String.fromCharCode(32);
    }
    return button;
}



//document.addEventListener("keydown", event => {});
