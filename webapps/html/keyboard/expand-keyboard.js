let ctrlLock = false;

function send(code){
    if(socket){
        socket.send(JSON.stringify({ type: "stdin", data: utf8_to_b64(code) }));
    }
}

$('#keyboard-esc').click(function(){
    let code = String.fromCharCode(27);
    send(code);
});

$('#keyboard-gang').click(function(){
    send("-");
});

$('#keyboard-up').click(function(){
    send(b64_to_utf8("G09B"));
});

$('#keyboard-paste').click(function(){
    alert(1);
    navigator.clipboard.readText()
        .then(text => {
            alert(2);
            send(text);
        })
        .catch(err => {
            alert(json.stringify(err));
            //console.error('Failed to read clipboard contents: ', err);
        });
});

$('#keyboard-sh-left').click(function(){
    send(b64_to_utf8("G1sxOzJE"));
});

$('#keyboard-tab').click(function(){
    let code = String.fromCharCode(9);
    send(code);
});

$('#keyboard-pie').click(function(){
    send(b64_to_utf8("Lw=="));
});

$('#keyboard-left').click(function(){
    send(b64_to_utf8("G09E"));
});

$('#keyboard-down').click(function(){
    send(b64_to_utf8("G09C"));
});

$('#keyboard-right').click(function(){
    send(b64_to_utf8("G09D"));
});

$('#keyboard-sh-right').click(function(){
    send(b64_to_utf8("G1sxOzJD"));
});


$('#keyboard-ctrl').click(function(){
    ctrlLock = true;
    lockCtrl();
});

function handleCtrl(button) {
    let code = ctrlToCharCode(button);
    if(socket){
        socket.send(JSON.stringify({ type: "stdin", data: utf8_to_b64(code) }));
    }
    releasedCtrl();
}

function lockCtrl() {
    let b = $('[data-skbtn="{controlleft}"]');
    b.addClass("ctrl-active");
    ctrlLock = true;
}
function releasedCtrl() {
    let b = $('[data-skbtn="{controlleft}"]');
    b.removeClass("ctrl-active");
    ctrlLock = false;
}

function ctrlToCharCode(button) {
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
    return button;
}

$(".expand-keyboard").mousedown(function(event){
    event.preventDefault();
});


