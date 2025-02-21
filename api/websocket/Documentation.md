# Broadcast Packet Documentation

This documentation outlines the structure and types of packets used for communication in the chess game.

## General Packet Structure

| **Field**         | **Size (bytes)** | **Description**                                                                 |
|--------------------|------------------|---------------------------------------------------------------------------------|
| **Packet Type**    | 1                | Identifies the type of message (e.g., move, game state, etc.).                  |
| **Body Length**    | 2                | Length of the body (number of bytes in the message data).                       |
| **Player ID**      | 1                | Identifies the player sending the message.                                      |
| **Body**           | Variable         | Contains the data specific to the packet type.                                  |

---

## Packet Types

### 1. **Move Packet (0x01)**

| **Field**               | **Size (bytes)** | **Description**                                                                 |
|--------------------------|------------------|---------------------------------------------------------------------------------|
| **From Square**          | 1                | Starting square (0-63, representing the chessboard index).                      |
| **To Square**            | 1                | Ending square (0-63, representing the target square).                           |
| **Move Metadata**        | 1                | Optional details like promotion type or special move.                           |

---

### 2. **Game State Packet (0x02)**

| **Field**               | **Size (bytes)** | **Description**                                                                 |
|--------------------------|------------------|---------------------------------------------------------------------------------|
| **Board State**          | 64               | 1 byte per square (each byte represents a piece or empty square).               |
| **Turn Indicator**       | 1                | Indicates which player’s turn it is.                                            |
| **Time Remaining**       | 4                | 2 bytes per player (time in seconds).                                           |

---

### 4. **Chat Packet (0x04)**

| **Field**               | **Size (bytes)** | **Description**                                                                 |
|--------------------------|------------------|---------------------------------------------------------------------------------|
| **Message Length**       | 1                | Length of the chat message.                                                     |
| **Chat Message**         | Variable         | Player-to-player chat message in UTF-8 encoding.                                |

---

### 5. **Legal Moves Request Packet (0x05)**

| **Field**               | **Size (bytes)** | **Description**                                                                 |
|--------------------------|------------------|---------------------------------------------------------------------------------|
| **Body**                | 0                | No additional data.                                                             |

---

### 6. **Legal Moves Response Packet (0x06)**

| **Field**               | **Size (bytes)** | **Description**                                                                 |
|--------------------------|------------------|---------------------------------------------------------------------------------|
| **Number of Moves (N)**  | 1                | Indicates how many legal moves are available.                                   |
| **Legal Moves**          | N                | List of legal moves (each represented by a byte for the target square).         |

---

### 7. **Game Over Packet (0x07)**

| **Field**               | **Size (bytes)** | **Description**                                                                 |
|--------------------------|------------------|---------------------------------------------------------------------------------|
| **Result Code**          | 1                | Indicates the reason for game termination (e.g., checkmate, stalemate).         |
| **Winner ID**            | 1                | Player identifier of the winner (0x00 for draw).                                |
| **Message (Optional)**   | Variable         | Reason or additional details about the game ending.                             |

---

### 9. **Resignation Packet (0x09)**

| **Field**               | **Size (bytes)** | **Description**                                                                 |
|--------------------------|------------------|---------------------------------------------------------------------------------|
| **Resigning Player ID**  | 1                | Indicates the player who resigned.                                              |

---

### 10. **Player Status Changed Packet (0x0A)**

| **Field**               | **Size (bytes)** | **Description**                                                                 |
|--------------------------|------------------|---------------------------------------------------------------------------------|
| **Player ID**            | 1                | Identifies the player whose status is updated.                                  |
| **Status Code**          | 1                | 0x00 for online, 0x01 for offline.                                              |

---

### Additional Packet Types

#### **Promotion Request Packet (0x0B)**

| **Field**           | **Size (bytes)** | **Description**                                                                 |
|----------------------|------------------|---------------------------------------------------------------------------------|
| **Pawn Position**    | 1                | Position of the pawn (0-63).                                                   |
| **New Piece Type**   | 1                | 0x01 for queen, 0x02 for rook, 0x03 for bishop, 0x04 for knight.               |

---

#### **Ping Packet (0x0C) & Ping Response Packet (0x0D)**

| **Field**           | **Size (bytes)** | **Description**                                                                 |
|----------------------|------------------|---------------------------------------------------------------------------------|
| **Body**            | 0                | No additional data.                                                             |

---

#### **Undo Move Request/Response Packet (0x0F/0x10)**

| **Field**           | **Size (bytes)** | **Description**                                                                 |
|----------------------|------------------|---------------------------------------------------------------------------------|
| **Result Code**      | 1 (Response)     | Indicates success or failure (response only).                                   |
| **Message**          | Variable         | Optional explanation (response only).                                          |

---

#### **Draw Offer/Response Packet (0x11/0x12)**

| **Field**           | **Size (bytes)** | **Description**                                                                 |
|----------------------|------------------|---------------------------------------------------------------------------------|
| **Player ID**        | 1                | Identifies the player offering or responding to the draw.                       |
| **Result Code**      | 1 (Response)     | Accept or reject (response only).                                              |

---

### **13. Ping Packet (0x0C)**

| **Field**   | **Size (bytes)** | **Description**                      |
|-------------|------------------|--------------------------------------|
| **Body**    | 0                | No additional data.                  |

---

### **14. Ping Response Packet (0x0D)**

| **Field**   | **Size (bytes)** | **Description**                      |
|-------------|------------------|--------------------------------------|
| **Body**    | 0                | Confirms server is alive.            |

---

### **15. Undo Move Request Packet (0x0F)**

| **Field**   | **Size (bytes)** | **Description**                      |
|-------------|------------------|--------------------------------------|
| **Body**    | 0                | No additional data.                  |

---

### **16. Undo Move Response Packet (0x10)**

| **Field**       | **Size (bytes)** | **Description**                      |
|------------------|------------------|--------------------------------------|
| **Result Code**  | 1                | Indicates success (0x01) or failure (0x00). |
| **Message**      | Variable         | Optional explanation of the result.  |

---

### **17. Draw Offer Packet (0x11)**

| **Field**       | **Size (bytes)** | **Description**                      |
|------------------|------------------|--------------------------------------|
| **Player ID**    | 1                | Identifies the player offering the draw. |

---

### **18. Draw Response Packet (0x12)**

| **Field**       | **Size (bytes)** | **Description**                      |
|------------------|------------------|--------------------------------------|
| **Result Code**  | 1                | Accept (0x01) or Reject (0x00).      |
| **Player ID**    | 1                | Identifies the player responding to the draw. |

---

### Summary of Packet Types

| **Packet Type**          | **Code** | **Description**                              |
|---------------------------|----------|----------------------------------------------|
| Move Packet              | 0x01     | Represents a chess move.                    |
| Game State Packet        | 0x02     | Represents the current state of the game.   |
| Chat Packet              | 0x04     | Sends player-to-player chat messages.       |
| Legal Moves Request      | 0x05     | Requests available legal moves.            |
| Legal Moves Response     | 0x06     | Responds with available legal moves.        |
| Game Over Packet         | 0x07     | Notifies that the game has ended.           |
| Resignation Packet       | 0x09     | Indicates that a player has resigned.       |
| Player Status Changed    | 0x0A     | Updates player status (online/offline).     |
| Promotion Request Packet | 0x0B     | Notifies a pawn promotion.                  |
| Ping Packet              | 0x0C     | Checks connectivity.                        |
| Ping Response Packet     | 0x0D     | Confirms server connectivity.               |
| Undo Move Request        | 0x0F     | Requests to undo the last move.             |
| Undo Move Response       | 0x10     | Responds to the undo request.               |
| Draw Offer Packet        | 0x11     | Offers a draw in the game.                  |
| Draw Response Packet     | 0x12     | Responds to a draw offer.                   |

---

