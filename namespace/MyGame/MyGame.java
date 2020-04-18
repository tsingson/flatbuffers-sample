// automatically generated by the FlatBuffers compiler, do not modify

package MyGame;

import java.nio.*;
import java.lang.*;
import java.util.*;
import com.google.flatbuffers.*;

@SuppressWarnings("unused")
public final class MyGame extends Table {
  public static void ValidateVersion() { Constants.FLATBUFFERS_1_12_0(); }
  public static MyGame getRootAsMyGame(ByteBuffer _bb) { return getRootAsMyGame(_bb, new MyGame()); }
  public static MyGame getRootAsMyGame(ByteBuffer _bb, MyGame obj) { _bb.order(ByteOrder.LITTLE_ENDIAN); return (obj.__assign(_bb.getInt(_bb.position()) + _bb.position(), _bb)); }
  public static boolean MyGameBufferHasIdentifier(ByteBuffer _bb) { return __has_identifier(_bb, "MGME"); }
  public void __init(int _i, ByteBuffer _bb) { __reset(_i, _bb); }
  public MyGame __assign(int _i, ByteBuffer _bb) { __init(_i, _bb); return this; }

  public byte CharactersType() { int o = __offset(4); return o != 0 ? bb.get(o + bb_pos) : 0; }
  public Table Characters(Table obj) { int o = __offset(6); return o != 0 ? __union(obj, o + bb_pos) : null; }

  public static int createMyGame(FlatBufferBuilder builder,
      byte Characters_type,
      int CharactersOffset) {
    builder.startTable(2);
    MyGame.addCharacters(builder, CharactersOffset);
    MyGame.addCharactersType(builder, Characters_type);
    return MyGame.endMyGame(builder);
  }

  public static void startMyGame(FlatBufferBuilder builder) { builder.startTable(2); }
  public static void addCharactersType(FlatBufferBuilder builder, byte CharactersType) { builder.addByte(0, CharactersType, 0); }
  public static void addCharacters(FlatBufferBuilder builder, int CharactersOffset) { builder.addOffset(1, CharactersOffset, 0); }
  public static int endMyGame(FlatBufferBuilder builder) {
    int o = builder.endTable();
    return o;
  }
  public static void finishMyGameBuffer(FlatBufferBuilder builder, int offset) { builder.finish(offset, "MGME"); }
  public static void finishSizePrefixedMyGameBuffer(FlatBufferBuilder builder, int offset) { builder.finishSizePrefixed(offset, "MGME"); }

  public static final class Vector extends BaseVector {
    public Vector __assign(int _vector, int _element_size, ByteBuffer _bb) { __reset(_vector, _element_size, _bb); return this; }

    public MyGame get(int j) { return get(new MyGame(), j); }
    public MyGame get(MyGame obj, int j) {  return obj.__assign(__indirect(__element(j), bb), bb); }
  }
}
