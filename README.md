waffle
===

`waffle` is a Lua execution environment, meant to be batteries-included.

The "twist" is that this isn't PUC Lua, it's backed by [gopher-lua][1]
and several preloaded modules.


Why?
===

Go's stdlib is comprehensive and I'm used to gopher-lua after using it in
several projects.

There's no design goal other than having fun.


Modules
===

Lots of modules come from `vadv/gopher-lua-libs` - many thanks to vadv and
contributors.

- [http][2]
- [json][3]
- [strings][4]
- [db][5]
- [tcp][6]
- [inspect][7]
- [time][8]
- [filepath][9]
- [cmd][10]
- [ioutil][11]
- [regexp][12]
- [plugin][13]
- [crypto][14]


[1]: https://github.com/yuin/gopher-lua
[2]: https://github.com/vadv/gopher-lua-libs/tree/master/http
[3]: https://github.com/vadv/gopher-lua-libs/tree/master/json
[4]: https://github.com/vadv/gopher-lua-libs/tree/master/strings
[5]: https://github.com/vadv/gopher-lua-libs/tree/master/db
[6]: https://github.com/vadv/gopher-lua-libs/tree/master/tcp
[7]: https://github.com/vadv/gopher-lua-libs/tree/master/inspect
[8]: https://github.com/vadv/gopher-lua-libs/tree/master/time
[9]: https://github.com/vadv/gopher-lua-libs/tree/master/filepath
[10]: https://github.com/vadv/gopher-lua-libs/tree/master/cmd
[11]: https://github.com/vadv/gopher-lua-libs/tree/master/ioutil
[12]: https://github.com/vadv/gopher-lua-libs/tree/master/regexp
[13]: https://github.com/vadv/gopher-lua-libs/tree/master/plugin
[14]: https://github.com/tengattack/gluacrypto

