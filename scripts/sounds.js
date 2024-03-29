var sounds = {
    "balbes" : {
        url : "sounds/balbes.mp3"
    },
    "zarezhu" : {
        url : "sounds/zarezhu.mp3"
    },
    "nahui" : {
        url : "sounds/nahui.mp3"
    },
    "boom" : {
        url : "sounds/boom.mp3"
    },
    "amogus" : {
        url : "sounds/amogus.mp3"
    },
    "midas" : {
        url : "sounds/midas.mp3"
    },
    "hah" : {
        url : "sounds/hah.mp3"
    },
    "volchok" : {
        url : "sounds/volchok.mp3"
    },
    "gachigasm" : {
        url : "sounds/gachigasm.mp3"
    },
    "mama" : {
        url : "sounds/mama.mp3"
    }
};


var soundContext = new AudioContext();

for (var sound in sounds) {
    loadSound(sound);
}

function loadSound(name) {
    var sound = sounds[name];
    var url = sound.url;
    var buffer = sound.buffer;
    var request = new XMLHttpRequest();

    request.open('GET', url, true);
    request.responseType = 'arraybuffer';

    request.onload = function() {
        soundContext.decodeAudioData(request.response, function(newBuffer) {
            sound.buffer = newBuffer;
        });
    }

    request.send();
}

function playSound(name, options) {
    var sound = sounds[name];
    var soundVolume = sounds[name].volume || 1;
    var buffer = sound.buffer;

    if (buffer) {
        var source = soundContext.createBufferSource();
        source.buffer = buffer;
        var volume = soundContext.createGain();

        if (options) {
            if (options.volume) {
                volume.gain.value = soundVolume * options.volume;
            }
        } else {
            volume.gain.value = soundVolume;
        }

        volume.connect(soundContext.destination);
        source.connect(volume);
        source.start(0);
    }
}

