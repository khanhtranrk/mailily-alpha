# mailily-alpha
mailily alpha

```
+-----------+       +---------+
| Requested |       | Justice |
| List      |       | Chain   |
+-----------+       +---------+

Create request -> send request -> add requested list

listen...
{
    requestId: 'xxxxx',
    safe_key: 'xxxxx'
    next_chain_token: 'xxxx',
    mes: '.....'
}

-> calculate next_chain_token(safe_token, safe_token, mes, previous_chain)
```

```
requestId: 'xxxx',
safe_key: 'xxxx',
previous_chain_token: 'xxxx',
previous_request_safe_key: 'xxxx'
```

```
+-----------+       +---------+
| Requested |       | Justice |
| List      |       | Chain   |
+-----------+       +---------+
```

```
problem 1: doing many request in one time
problem 2: Justice Book
```

Justice Book (mavis[hidden_key] - signima[hidden_key])
- requestId | chainKey
- requestId | chainKey
- requestId | chainKey
- requestId | chainKey
- requestId | chainKey
- requestId | chainKey

Mavis -> request -> send request --(done)-> add to Reqeusted List
Signima -> receive request -> request list -> status: waiting

doing request 1
    task new onetime [time] [mail account list]

task doing task_id -> requestid reponse -> sender -> requestId -> send



mavis -> (sender: mavis, receiver: signima, mes: ...)
-> signima -> 

In Request List
request_id - Keep

request_id: (sender: mavis, receiver: signima, mes: ...)

Out Request List
request 1 - Keep
request 2 - Done

mavis -> signima -> task
Justice Book

mavis +1
signima +2
task +1

mavis +1
signima +2
task


In Request
Request -> x

Out Request

```
Sender: signima
Receiver: mavis
safe_key
next_chain_token
```

```
Type: request | response
```

 In Request | Out Request
 -----------|------------
 xxxx       | xxxx
 yyyy       | yyyy
 zzzz       | zzzz

```
chan msgs <- contactGate.Listen()

// handle After

// response After

// received Message

// sent Message

// nclcnqdml,dnnqn

// nvhnltkmcvmnnct

// vtpcn,kcndtpck

// tnhktntdpcnntn

// tvttntnvcd,inctsddnltbvstdhtvcnnbg

// nmqtctvllmcnctndbsatcbt

handleReqeust(m) { 
    m := saveMessage(m, 0) // message, status
    context, err := exec(m.Content)
    m = buildMessageFromContext(context, m)
    context, err := taistra.SendTo(m)
    err := justice.CouldNotConnectTaistra()....
    if err {
        responseAfter(m)
    } else {
        
    }
}

for msg in msgs {
    t, m = convertToMessage(msg)
    err := CheckingJusticeInfo()
    if err != nil {
        err := justice.CouldNotConnectTaistra()
        pingToTaistra(...) // ... //\\
        handleAfter(m)
    }

    if t == 1 {
        handleReqeust(m)
    } else if t == 2 {
        handleResponse
    } else {
        handleError(...)
    }
}

func exec1(m []byte) {
    seekM = m

    while len(seek) != 0 || seekM[0] != 0 (end) {
        if byte[1] == 1 {
            context {status, mess}, err := DoSomething()
            if mess, err {
            }
        } else if byte[1] == 2 {
            seek, err := exec1(...)
        } else {
            return nil, err
        }

        return err if err

        seekM = seek
    }

    return nil
}

func exec(m []byte) {
    seekM = m

    while len(seek) != 0 || seekM[0] != 0 (end) {
        if byte[1] == 1 {
            seek, err := exec1(...)
        } else if byte[1] == 2 {
            seek, err := exec1(...)
        } else {
            return nil, err
        }

        return err if err

        seekM = seek
    }

    return nil
}

// The human life

func (me *Human) RunLife() {
    go me.MakeMoney(nil, nil)

    for {
        select {
        case m <- me.money:
            me.balance += m.value
        case t <- me.tired:
            if t {
                me.tired <- False
            }
        }
        case d <- me.died:
            Close(me.money)
            Close(me.balance)
            close(me.life)
            log.Fatal("The end")
    }
}

func (me *Human) MakeMoney(opportunity *Opportunity = nil, money Money = nil) {
    for {
        if not opportunity {
            opportunity = me.FindOpportunity()
        }

        go me.MakeMoneyFromOpportunity(opportunity, money)
    }
}

func (me *Human) MakeMoneyFromOpportunity(opportunity *Opportunity, money Money) {
    me.money <- opportunity.MakeMoney(money)
}
```
