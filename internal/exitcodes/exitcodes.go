package exitcodes

const (
	// ========================================================================.
	// Restricting user-defined exit codes to the range 64 - 113 (in addition to
	// 0, for success) allows for 50 exit codes. The range 114-125 *could* be
	// used (adding another ten codes) but DO NOT GO BEYOND 126!
	// ========================================================================.

	// ========================================================================.
	// Exit Codes With Special Meanings
	// ------------------------------------------------------------------------.

	Ok int = 0

	// GeneralError
	// Catchall for general errors: Miscellaneous errors, such as 'divide by
	// zero' and other impermissible operations.
	GeneralError int = 1

	// MisuseOfShellBuiltins Misuse of shell builtins: Missing keyword or command,
	// or permission problem (and diff return code on a failed binary file
	// comparison).
	MisuseOfShellBuiltins int = 2
	// ========================================================================.

	// ========================================================================.
	// Reserved codes (range 3-63)
	// ========================================================================.

	// ========================================================================.
	// Generic errors (60 range)
	// ------------------------------------------------------------------------.

	UnknownErrorOccurred       int = 64
	NotEnoughParameters        int = 65
	InvalidParameter           int = 66
	DependencyMissing          int = 67
	EnvironmentVariableMissing int = 68
	CouldNotDownload           int = 69
	// ========================================================================.

	// ========================================================================.
	// File and folder errors (70 range)
	// ------------------------------------------------------------------------.

	CouldNotFindFile        int = 70
	CouldNotCreateFile      int = 71
	CouldNotUpdateFile      int = 72
	CouldNotDeleteFile      int = 73
	CouldNotMoveFile        int = 74
	CouldNotFindDirectory   int = 75
	CouldNotCreateDirectory int = 76
	CouldNotUpdateDirectory int = 77
	CouldNotDeleteDirectory int = 78
	CouldNotMoveDirectory   int = 79
	// ========================================================================.

	// ========================================================================.
	// Other errors (80 range)
	// ------------------------------------------------------------------------.

	CouldNotFind       int = 80
	CouldNotCreate     int = 81
	CouldNotUpdate     int = 82
	CouldNotDelete     int = 83
	CouldNotMove       int = 84
	CouldNotRead       int = 85
	CouldNotReadFile   int = 86
	CouldNotReadFolder int = 87
	NotCorrectType     int = 88
	NotSupported       int = 89
	// ========================================================================.

	// ========================================================================.
	// Credential and validation errors (90 range)
	// ------------------------------------------------------------------------.

	ValidationFailed   int = 90
	CredentialsMissing int = 91
	CredentialsInvalid int = 92
	// ========================================================================.

	// ========================================================================.
	// Application specific errors (93-113 range)
	// These should be reserved for program specific errors
	// ========================================================================.

	// ========================================================================.
	// Reserved codes (range 114-125)
	// ========================================================================.

	// ========================================================================.
	// Exit Codes With Special Meanings
	// ------------------------------------------------------------------------.

	// CannotExecuteCommand Command invoked cannot execute, permission problem
	// or command is not an executable.
	CannotExecuteCommand int = 126
	// CommandNotFound "command not found", possible problem with $PATH or a typo.
	CommandNotFound int = 127
	// InvalidArgument Invalid argument to exit, exit takes only integer args in
	// the range 0-255.
	InvalidArgument int = 128
	// ========================================================================.

	// ========================================================================.
	// Fatal error signal (Range 129-192) from POSIX signals
	// For signal n:  exit code = 128+n
	// ------------------------------------------------------------------------.

	SigHup int = 129 // signal 1: (128+1)=129

	// SigInt Script terminated by Control-C, Control-C is fatal error signal 2.
	SigInt     int = 130 // signal 2: (128+2)=130
	SigQuit    int = 131 // signal 3: (128+3)=131
	SigIll     int = 132 // signal 4: (128+4)=132
	SigTrap    int = 133 // signal 5: (128+5)=133
	SigAbrt    int = 134 // signal 6: (128+6)=134
	SigBus     int = 135 // signal 7: (128+7)=135
	SigFpe     int = 136 // signal 8: (128+8)=136
	SigKill    int = 137 // signal 9: (128+9)=137
	SigUsr1    int = 138 // signal 10: (128+10)=138
	SigSegv    int = 139 // signal 11: (128+11)=139
	SigUsr2    int = 140 // signal 12: (128+12)=140
	SigPipe    int = 141 // signal 13: (128+13)=141
	SigAlrm    int = 142 // signal 14: (128+14)=142
	SigTerm    int = 143 // signal 15: (128+15)=143
	SigStkflt  int = 144 // signal 16: (128+16)=144
	SigChld    int = 145 // signal 17: (128+17)=145
	SigCont    int = 146 // signal 18: (128+18)=146
	SigStop    int = 147 // signal 19: (128+19)=147
	SigTstp    int = 148 // signal 20: (128+20)=148
	SigTtin    int = 149 // signal 21: (128+21)=149
	SigTtou    int = 150 // signal 22: (128+22)=150
	SigUrg     int = 151 // signal 23: (128+23)=151
	SigXcpu    int = 152 // signal 24: (128+24)=152
	SigXfsz    int = 153 // signal 25: (128+25)=153
	SigVtalrm  int = 154 // signal 26: (128+26)=154
	SigProf    int = 155 // signal 27: (128+27)=155
	SigWinch   int = 156 // signal 28: (128+28)=156
	SigIo      int = 157 // signal 29: (128+29)=157
	SigPwr     int = 158 // signal 30: (128+30)=158
	SigSys     int = 159 // signal 31: (128+31)=159
	SigUnknown int = 160 // signal 32: (128+32)=160
	// SigUnknown int = 161 // signal 33: (128+33)=161.

	SigRtmin        int = 162 // signal 34: (128+34)=162
	SigRtminPlus1   int = 163 // signal 35: (128+35)=163
	SigRtminPlus2   int = 164 // signal 36: (128+36)=164
	SigRtminPlus3   int = 165 // signal 37: (128+37)=165
	SigRtminPlus4   int = 166 // signal 38: (128+38)=166
	SigRtminPlus5   int = 167 // signal 39: (128+39)=167
	SigRtminPlus6   int = 168 // signal 40: (128+40)=168
	SigRtminPlus7   int = 169 // signal 41: (128+41)=169
	SigRtminPlus8   int = 170 // signal 42: (128+42)=170
	SigRtminPlus9   int = 171 // signal 43: (128+43)=171
	SigRtminPlus10  int = 172 // signal 44: (128+44)=172
	SigRtminPlus11  int = 173 // signal 45: (128+45)=173
	SigRtminPlus12  int = 174 // signal 46: (128+46)=174
	SigRtminPlus13  int = 175 // signal 47: (128+47)=175
	SigRtminPlus14  int = 176 // signal 48: (128+48)=176
	SigRtminPlus15  int = 177 // signal 49: (128+49)=177
	SigRtmaxMinus14 int = 178 // signal 50: (128+50)=178
	SigRtmaxMinus13 int = 179 // signal 51: (128+51)=179
	SigRtmaxMinus12 int = 180 // signal 52: (128+52)=180
	SigRtmaxMinus11 int = 181 // signal 53: (128+53)=181
	SigRtmaxMinus10 int = 182 // signal 54: (128+54)=182
	SigRtmaxMinus9  int = 183 // signal 55: (128+55)=183
	SigRtmaxMinus8  int = 184 // signal 56: (128+56)=184
	SigRtmaxMinus7  int = 185 // signal 57: (128+57)=185
	SigRtmaxMinus6  int = 186 // signal 58: (128+58)=186
	SigRtmaxMinus5  int = 187 // signal 59: (128+59)=187
	SigRtmaxMinus4  int = 188 // signal 60: (128+60)=188
	SigRtmaxMinus3  int = 189 // signal 61: (128+61)=189
	SigRtmaxMinus2  int = 190 // signal 62: (128+62)=190
	SigRtmaxMinus1  int = 191 // signal 63: (128+63)=191
	SigRtmax        int = 192 // signal 64: (128+64)=192
)
